package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// コメントアウトは律儀に読み飛ばしてくれるらしい
// go run main.go url [url ...]
func main() {
	for i := 1; i < len(os.Args); i++ {
		url := os.Args[i]
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("url:", url)
		fmt.Println("words:", words)
		fmt.Println("images:", images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// 返り値省略、確かにみにくい
func countWordsAndImages(n *html.Node) (words, images int) {
	// 不可視の要素はskip
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
		if n.Data == "img" {
			images++
		}
	}
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	return
}
