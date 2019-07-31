package main

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	// 正しく検出できているかのテスト
	var tests = []struct {
		s    string
		want string
	}{
		{"hello world $くれ😣なゐ $$$$$ $#a# $たばた $しんぶんし hello $", "hello world くれ😣なゐ $$$$ #a# たばた しんぶんし hello "},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("expand(%s, func())", test.s)
		got := expand(test.s, func(s string) string {
			return s
		})

		if got != test.want {
			t.Errorf("%s = \n%q, want \n%q", descr, got, test.want)
		}
	}
}
