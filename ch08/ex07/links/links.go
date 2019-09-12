package links

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func URL2Filepath(uri *url.URL) string {
	path := filepath.Join("./", uri.Path)
	if len(uri.Path) == 0 || uri.Path[len(uri.Path)-1] == '/' {
		path = filepath.Join(path, "index.html")
	}
	if !strings.HasSuffix(path, ".html") {
		path += ".html"
	}
	return path
}

// isMirror: falseにすると保存していないパスにはミラーリングしなくなる
func Extract(out io.Writer, uri string, isMirror bool, rootPath string) ([]string, error) {
	baseURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	depth := 0
	res := []string{}
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Fprintf(out, "%*s<%s", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				// linkの時
				if n.Data == "a" && attr.Key == "href" {
					link, err := resp.Request.URL.Parse(attr.Val)
					if err != nil {
						continue // ignore bad URLs
					}
					res = append(res, link.String())

					if link.Hostname() == baseURL.Hostname() {
						localPath := URL2Filepath(link)
						if isMirror { // ミラーリングする
							fmt.Printf("%s -> %s\n", link, localPath) // debug log output
							fmt.Fprintf(out, " %s=%q", attr.Key, localPath)
						} else { //ミラーリングをしない
							_, err := os.Stat(filepath.Join(rootPath, localPath))
							fmt.Printf("%s -> %s\n", link, localPath) // debug log output
							if err == nil {                           // ファイルが存在している
								fmt.Printf("%s -> %s\n", link, localPath) // debug log output
								fmt.Fprintf(out, " %s=%q", attr.Key, localPath)
							} else {
								fmt.Printf("%s -> %s\n", link, link.String()) // debug log output
								fmt.Fprintf(out, " %s=%q", attr.Key, link.String())
							}
						}
					} else {
						fmt.Fprintf(out, " %s=%q", attr.Key, link.String())
					}
				} else {
					fmt.Fprintf(out, " %s=%q", attr.Key, attr.Val)
				}
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

	endElement := func(n *html.Node) {
		if n.FirstChild != nil && n.Type == html.ElementNode {
			depth--
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}

	}

	forEachNode(doc, startElement, endElement)
	return res, nil
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
