package main

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		rs   []rune
		want string
	}{
		{[]rune{' ', '\t', 'H', 'e', 'l', 'l', 'o', '\t', ' ', '　', '\t', '世', '界', '\r', '\v'}, " Hello 世界 "},
		{[]rune{'H', 'e', 'l', 'l', 'o', '\t', ' ', '　', '\t', '世', '界'}, "Hello 世界"},
		{[]rune{'H', 'e', 'l', 'l', 'o', '世', '界'}, "Hello世界"},
	}

	for _, test := range tests {
		bs := []byte(string(test.rs))
		descr := fmt.Sprintf("compressSpace(%v)", bs)
		got := string(compressSpace(bs))
		if got != test.want {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}
