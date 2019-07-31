package main

import (
	"fmt"
	"io"
	"os"

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

	res := ElementByID(doc, "kurenaif")

	// print output
	fmt.Println(res.Type, html.ElementNode)

	if res.Type == html.ElementNode {
		fmt.Fprintf(out, "%*s<%s", depth*2, "", res.Data)
		for _, attr := range res.Attr {
			fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
		}
		fmt.Fprintf(out, ">")
	}

	return nil
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

var depth int

// func startElement(n *html.Node) bool {
// 	if n.Type == html.ElementNode {
// 		id, ok := n.Attr[attr.Key]
// 		for _, attr := range n.Attr {
// 			fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
// 		}
// 	}
// 	if n.Type == html.ElementNode {
// 		fmt.Fprintf(out, "%*s<%s", depth*2, "", n.Data)
// 		for _, attr := range n.Attr {
// 			fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
// 		}
// 		if n.FirstChild != nil {
// 			fmt.Fprintf(out, ">")
// 			depth++
// 		} else {
// 			fmt.Fprintf(out, "/>")
// 		}
// 		fmt.Fprintf(out, "\n")
// 	}

// 	if n.Type == html.TextNode {
// 		text := n.Data
// 		text = strings.TrimSpace(text)
// 		if text != "" {
// 			fmt.Fprintf(out, "%*s%s\n", depth*2, "", text)
// 		}
// 	}

// 	if n.Type == html.CommentNode {
// 		text := n.Data
// 		text = strings.TrimSpace(text)
// 		if text != "" {
// 			fmt.Fprintf(out, "%*s<!-- %s -->\n", depth*2, "", text)
// 		}
// 	}
// }

// func endElement(n *html.Node) bool {
// 	if n.FirstChild != nil && n.Type == html.ElementNode {
// 		depth--
// 		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }
