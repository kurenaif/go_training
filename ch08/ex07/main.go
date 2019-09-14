package main

import (
	"flag"
	"go_training/ch08/ex07/links"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func crawl(uri string, depth int, saveDir string, isMirror bool) ([]string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	savePath := links.URL2Filepath(u)
	savePath = filepath.Join(saveDir, savePath)
	os.MkdirAll(filepath.Dir(savePath), 0775)
	file, err := os.Create(savePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	list, err := links.Extract(file, uri, isMirror, saveDir)
	if err != nil {
		return nil, err
	}
	return list, nil
}

type LinksCnt struct {
	links []string
	cnt   int
}

type LinkCnt struct {
	link     string
	cnt      int
	hostname string
	isMirror bool
}

//!+
func main() {
	var depth = flag.Int("depth", -1, "link depth")
	flag.Parse()
	if depth == nil || *depth < 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	worklist := make(chan LinksCnt)   // lists of URLs, may have duplicates
	unseenLinks := make(chan LinkCnt) // de-duplicated URLs

	// Add command-line arguments to worklist.
	n := 0
	n++
	go func() { worklist <- LinksCnt{flag.Args(), 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks, err := crawl(link.link, link.cnt, link.hostname, link.isMirror)
				if err != nil {
					log.Print(err)
					continue
				}
				go func(cnt int) { worklist <- LinksCnt{foundLinks, cnt + 1} }(link.cnt)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	baseHostnames := map[string]bool{}

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.links {
			u, err := url.Parse(link)
			if err != nil {
				log.Print(err)
				continue
			}
			if !seen[link] {
				seen[link] = true
				if list.cnt == 0 {
					baseHostnames[u.Hostname()] = true
				} else {
					if _, ok := baseHostnames[u.Hostname()]; !ok {
						continue
					}
				}
				if list.cnt < *depth { // depth 3 では、3つめは見る
					n++
					unseenLinks <- LinkCnt{link, list.cnt, u.Hostname(), (false)} //次終了する => ミラーリングを終了する
				}
			}
		}
	}
}

//!-
