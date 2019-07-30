package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for nodeType, num := range visit(nil, doc) {
		fmt.Println(nodeType, num)
	}
}

func visit(cnt map[string]int, n *html.Node) map[string]int {
	if cnt == nil {
		cnt = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		cnt[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cnt = visit(cnt, c)
	}
	return cnt
}
