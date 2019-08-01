package main

import (
	"testing"
)

func TestTSort(t *testing.T) {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},

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

	for i := 0; i < 10; i++ { //結果が不定なので何回かやる
		order := topoSort(prereqs)
		check := make(map[string]bool)
		for _, item := range order {
			// 履修するための条件チェック
			items, ok := prereqs[item]
			if ok {
				for _, req := range items {
					_, ok := check[req]
					if !ok {
						t.Errorf("Prerequisite %s for receiving class %s is not satisfied.", req, item)
					}
				}
			}
			// 履修した
			check[item] = true
		}
	}
}
