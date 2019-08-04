package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

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

func crawl(pathString string) []string {
	fmt.Println(pathString)
	list := []string{}

	info, err := os.Stat(pathString)
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		fileInfos, err := ioutil.ReadDir(pathString)
		if err != nil {
			log.Print(err)
		}
		for _, fileInfo := range fileInfos {
			newPath := path.Join(pathString, fileInfo.Name())
			list = append(list, newPath)
		}
	}

	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
