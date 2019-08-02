package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type PrePostElement struct {
	pre  func(*html.Node)
	post func(*html.Node)
}

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startEndElement())

	return nil
}

func forEachNode(n *html.Node, startEnd PrePostElement) {
	prePost := startEnd

	if prePost.pre != nil {
		prePost.pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, startEnd)
	}

	if prePost.post != nil {
		prePost.post(n)
	}
}

func startEndElement() PrePostElement {
	var depth int

	return PrePostElement{
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		},
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		},
	}
}
