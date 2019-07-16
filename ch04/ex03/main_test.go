package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		a    [6]int
		want [6]int
	}{
		{[6]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", test.a)
		reverse(&test.a)
		got := test.a
		if got != test.want {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}
