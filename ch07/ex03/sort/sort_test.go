// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		values []int
		want   string
	}{
		{[]int{1, 2, 3, 4, 5}, "[1 2 3 4 5]"},
		{[]int{5, 1, 3, 2, 4}, "[1 2 3 4 5]"},
		{[]int{}, "[]"},
		{[]int{0}, "[0]"},
		{[]int{-5, 5}, "[-5 5]"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("values = %v; tree.String()", test.values)
		var root *tree
		for _, v := range test.values {
			root = add(root, v)
		}
		got := root.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}
