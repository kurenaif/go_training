package main

import (
	"bytes"
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
		want    string
	}{
		{"<html><head><title>title</title></head><body><a href=\"link1\">hello</a><p>world</p></body></html>", "title\nhello\nworld\n"},                                                                       // title
		{"<html><head></head><body><a href=\"link1\"></a><a href=\"link2\"></a></body></html>", ""},                                                                                                           // none
		{"<html><head><title>title</title><style>body{backgruond-color: #00ffff}</style></head><body><a href=\"link1\">hello</a><p>world</p><script>alert(\"hello\")</body></html>", "title\nhello\nworld\n"}, // script, style skip
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(%s)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		out = new(bytes.Buffer)
		visit(doc)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
