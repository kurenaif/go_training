package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		lhs  string
		rhs  string
		want bool
	}{
		{"aabbcc", "abcabc", true},
		{"aaabbcc", "abcabc", false},
		{"aabbcc", "aabcabc", false},
		{"bbcc", "aabcbc", false},
		{"aabbcc", "bcbc", false},
		{"abc世界世界abc", "aabbcc世世界界", true},
		{"abc世界世界abc", "aabbcc世界界", false},
		{"aAあ亞アд①😥", "あ亞アд①😥aA", true},
		{"aAあ亞アд①😥", "あ亞ア①😥aA", false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("isAnagram(%q, %q)", test.lhs, test.rhs)
		got := isAnagram(test.lhs, test.rhs)
		if got != test.want {
			t.Errorf("%s = %t, want %t", descr, got, test.want)
		}
	}
}
