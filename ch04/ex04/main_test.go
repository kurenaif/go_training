package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	var tests = []struct {
		lhs  int
		rhs  int
		want int
	}{
		{65537, 12345, 1},
		{32, 20, 4},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("gcd(%d, %d)", test.lhs, test.rhs)
		got := gcd(test.lhs, test.rhs)
		if got != test.want {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}

func equal(x, y []int) bool {
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

func TestRotate(t *testing.T) {
	var tests = []struct {
		s      []int
		offset int
		want   []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 0, []int{0, 1, 2, 3, 4, 5}},
		{[]int{0, 1, 2, 3, 4, 5}, 1, []int{5, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4, 5}, 7, []int{5, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{4, 5, 0, 1, 2, 3}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("rotate(%v, %d)", test.s, test.offset)
		rotate(test.s, test.offset)
		got := test.s
		if !equal(got, test.want) {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}
