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

type count struct {
	words  int
	images int
}

func TestVisit(t *testing.T) {
	var tests = []struct {
		content string
		want    count
	}{
		{"<p>hello world</p>", count{2, 0}},
		{"<div><p>hello world</p></div>", count{2, 0}},
		{"<div><p>hello world</p><div><p>kijin seija</p></div></div>", count{4, 0}},
		{"<div><p>hello world</p><div><p>kijin seija</p><img src=\"seija.png\"/></div></div>", count{4, 1}},
		{"<img src=\"hello.jpg\"/>", count{0, 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(nil, %s)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		words, images := countWordsAndImages(doc)
		got := count{words, images}
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
