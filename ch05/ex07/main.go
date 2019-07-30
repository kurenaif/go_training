package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	// for _, url := range os.Args[1:] {
	// 	outline(url)
	// }
	outline("hello")
}

func outline(url string) error {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
		}
		if n.FirstChild != nil {
			fmt.Fprintf(out, ">")
			depth++
		} else {
			fmt.Fprintf(out, "/>")
		}
		fmt.Fprintf(out, "\n")
	}

	if n.Type == html.TextNode {
		text := n.Data
		text = strings.TrimSpace(text)
		if text != "" {
			fmt.Fprintf(out, "%*s%s\n", depth*2, "", text)
		}
	}

	if n.Type == html.CommentNode {
		text := n.Data
		text = strings.TrimSpace(text)
		if text != "" {
			fmt.Fprintf(out, "%*s<!-- %s -->\n", depth*2, "", text)
		}
	}
}

func endElement(n *html.Node) {
	if n.FirstChild != nil && n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
