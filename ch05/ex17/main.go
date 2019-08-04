package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println("--------------------img--------------------")
		nodes, err := URLElementByTagName(url, "img")
		if err != nil {
			continue
		}
		for _, node := range nodes {
			fmt.Println(node)
		}
		fmt.Println("--------------------h1,2,3,4--------------------")
		nodes, err = URLElementByTagName(url, "h1", "h2", "h3", "h4")
		if err != nil {
			continue
		}
		for _, node := range nodes {
			fmt.Println(node)
		}
	}
}

func URLElementByTagName(url string, name ...string) ([]*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	nodes := ElementByTagName(doc, name...)

	return nodes, nil
}

func ElementByTagName(doc *html.Node, name ...string) (res []*html.Node) {
	nameSet := make(map[string]bool)
	for _, na := range name {
		nameSet[na] = true
	}

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {
			if _, ok := nameSet[n.Data]; ok {
				res = append(res, n)
			}
		}
	})

	return res
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	pre(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}
