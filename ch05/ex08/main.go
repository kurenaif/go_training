package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url id", os.Args[0])
		return
	}

	ok, err := findElementById(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		fmt.Fprintln(os.Stderr, "id not found")
	}
}

func findElementById(url string, id string) (bool, error) { // 見つかったかどうかを返す
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return false, err
	}

	res := ElementByID(doc, id)
	if res == nil {
		return false, nil
	}

	// output html Node
	if res.Type == html.ElementNode {
		fmt.Fprintf(out, "<%s", res.Data)
		for _, attr := range res.Attr {
			fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
		}
		fmt.Fprintf(out, ">")
	}

	return true, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) { // is skip => true
	if pre != nil {
		if pre(n) { // is skip => true
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if post(n) { // is skip => true
			return
		}
	}

	return
}

func ElementByID(doc *html.Node, id string) (idNode *html.Node) {
	isFinish := false

	forEachNode(doc, func(n *html.Node) bool { //pre
		if isFinish {
			return true
		}
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					idNode = n
					isFinish = true
					return true
				}
			}
		}
		return false
	}, func(*html.Node) bool { // post 特に何もしない
		return false
	})
	return
}
