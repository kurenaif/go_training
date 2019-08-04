package main

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		sep  string
		strs []string
		want string
	}{
		{",,,", []string{"a", "b", "c"}, "a,,,b,,,c"},
		{"", []string{"abc", "def", "ghi"}, "abcdefghi"},
		{",", []string{"a"}, "a"},
		{",", []string{""}, ""},
		{"", []string{""}, ""},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("join(%s, %v)", test.sep, test.strs)
		got := join(test.sep, test.strs...)
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}
