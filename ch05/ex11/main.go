package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// ref) https://ja.wikipedia.org/wiki/%E3%83%88%E3%83%9D%E3%83%AD%E3%82%B8%E3%82%AB%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88
func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]int) // 1: 一時的な印, 2: 恒久的な印
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if seen[item] == 1 {
				return fmt.Errorf("graph contain closed circuit processing %s", item)
			}
			if seen[item] == 0 {
				seen[item] = 1
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				seen[item] = 2
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}
