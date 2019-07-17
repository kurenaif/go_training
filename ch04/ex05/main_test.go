package main

import (
	"fmt"
	"testing"
)

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestUnique(t *testing.T) {
	var tests = []struct {
		ss   []string
		want []string
	}{
		{[]string{"a", "a", "b", "b", "a", "b"}, []string{"a", "b", "a", "b"}},
		{[]string{"a", "b", "a", "b"}, []string{"a", "b", "a", "b"}},
		{[]string{"a", "a", "a", "a"}, []string{"a"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("unique(%v)", test.ss)
		got := unique(test.ss)
		if !equal(got, test.want) {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}
