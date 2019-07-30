package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func isSame(lhs map[string]int, rhs map[string]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, lhsValue := range lhs {
		if lhsValue != rhs[k] {
			return false
		}
	}
	return true
}

func TestVisit(t *testing.T) {
	var tests = []struct {
		content string
		want    map[string]int
	}{
		{"<html><head></head><body><a href=\"link1\"></a></body></html>", map[string]int{"a": 1, "html": 1, "body": 1, "head": 1}},
		{"<html><head></head><body><a href=\"link1\"></a><a href=\"link2\"></a></body></html>", map[string]int{"a": 2, "html": 1, "body": 1, "head": 1}},
		{"<html><head></head><body><a href=\"link1\"></a><div><a href=\"link2\"></a></div></body></html>", map[string]int{"a": 2, "div": 1, "html": 1, "body": 1, "head": 1}},
		{"<html><head></head><body><a href=\"link1\"></a><div><div><a href=\"link2\"></a></div><a href=\"link3\"></a></div></body></html>", map[string]int{"a": 3, "div": 2, "html": 1, "body": 1, "head": 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(nil, %s)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		got := visit(nil, doc)
		if !isSame(got, test.want) {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
