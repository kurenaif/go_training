// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

var whiteHostSet = make(map[string]bool)

func saveUrl(url string, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(savePath), 0775)

	if err := ioutil.WriteFile(savePath, body, 0664); err != nil {
		return err
	}

	return nil
}

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlString string) []string {
	list, err := links.Extract(urlString)
	if err != nil {
		log.Print(err)
		return list
	}
	u, err := url.Parse(urlString)
	if err != nil {
		log.Print(err)
		return list
	}
	if _, ok := whiteHostSet[u.Host]; ok {
		if u.Path == "" {
			u.Path = "/"
		}
		path := u.Host + u.Path
		log.Println("saving... ", path)
		if strings.HasSuffix(path, "/") {
			path += "index.html"
		}
		fmt.Println(path)
		saveUrl(urlString, path)
	}
	return list
}

//!-crawl

//!+main
func main() {
	for _, urlString := range os.Args[1:] {
		url, err := url.Parse(urlString)
		if err != nil {
			log.Print(err)
		}
		whiteHostSet[url.Host] = true
	}
	breadthFirst(crawl, os.Args[1:])
}

//!-main
