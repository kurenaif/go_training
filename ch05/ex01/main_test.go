package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func isSame(lhs []string, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func TestVisit(t *testing.T) {
	var tests = []struct {
		content string
		want    []string
	}{
		{"<html><head></head><body><a href=\"link1\"></a></body></html>", []string{"link1"}},
		{"<html><head></head><body><a href=\"link1\"></a><a href=\"link2\"></a></body></html>", []string{"link1", "link2"}},
		{"<html><head></head><body><a href=\"link1\"></a><div><a href=\"link2\"></a></div></body></html>", []string{"link1", "link2"}},
		{"<html><head></head><body><a href=\"link1\"></a><div><div><a href=\"link2\"></a></div><a href=\"link3\"></a></div></body></html>", []string{"link1", "link2", "link3"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(nil, %s)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		got := visit(nil, doc)
		if !isSame(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
