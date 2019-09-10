// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"go_training/ch08/ex06/links"
	"log"
	"os"
)

func crawl(url string, depth int) []string {
	fmt.Println(depth, url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type LinksCnt struct {
	links []string
	cnt   int
}

type LinkCnt struct {
	link string
	cnt  int
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
				foundLinks := crawl(link.link, link.cnt)
				go func(cnt int) { worklist <- LinksCnt{foundLinks, cnt + 1} }(link.cnt)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				if list.cnt < *depth { // depth 3 では、3つめは見る
					n++
					unseenLinks <- LinkCnt{link, list.cnt}
				}
			}
		}
	}
}

//!-
