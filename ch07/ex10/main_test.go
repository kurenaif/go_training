package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    []int
		want bool
	}{
		{[]int{0, 1, 2, 1, 0}, true},
		{[]int{0, 1, 2, 1, 2}, false},
		{[]int{0, 1, 2, 2, 1, 0}, true},
		{[]int{0}, true},
		{[]int{}, true},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IsPalindrome(%v)", test.s)
		got := IsPalindrome(sort.IntSlice(test.s))
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
