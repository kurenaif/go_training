package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567890", "1,234,567,890"},
		{"234567890", "234,567,890"},
		{"34567890", "34,567,890"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("comma(%q)", test.s)
		got := comma(test.s)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
