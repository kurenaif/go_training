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

	order, err := topoSort(prereqs)
	if err != nil {
		t.Errorf("this graph has not contain closed circruit. but topoSort output this error: %s", err)
	}
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

func TestClosedTSort(t *testing.T) {
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

	order, err := topoSort(prereqs)
	if err == nil {
		t.Errorf("this graph has contain closed circruit. but topoSort didn't output error")
	}
	if order != nil {
		t.Errorf("order want: nil, but got: %s", order)
	}

}
