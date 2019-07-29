// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
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
	for nodeType, num := range visit(make(map[string]int), doc) {
		fmt.Println(nodeType, num)
	}
}

func visit(cnt map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		cnt[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cnt = visit(cnt, c)
	}
	return cnt
}
