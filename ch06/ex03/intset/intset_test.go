package intset

import (
	"fmt"
	"testing"
)

// Stringが常に昇順であることを仮定した比較方法
func isSame(lhs IntSet, rhs IntSet) bool {
	return lhs.String() == rhs.String()
}

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		s    []int // 入力する値
		t    []int
		want []int // 最終的に求める状態
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{0, 1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6}, []int{0}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6, 123456}, []int{0}},
		{[]int{0, 1, 2, 3, 123456}, []int{0, 4, 5, 6}, []int{0}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.s {
			s.Add(value)
		}
		var tSet IntSet
		for _, value := range test.t {
			tSet.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).InterSectWith(%v)", s.String(), tSet.String())
		s.IntersectWith(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		s    []int // 入力する値
		t    []int
		want []int // 最終的に求める状態
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{0, 1, 2, 3}, []int{4, 5, 6}, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6}, []int{1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6, 123456}, []int{1, 2, 3}},
		{[]int{0, 1, 2, 3, 123456}, []int{0, 4, 5, 6}, []int{1, 2, 3, 123456}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.s {
			s.Add(value)
		}
		var tSet IntSet
		for _, value := range test.t {
			tSet.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).DifferenceWith(%v)", s.String(), tSet.String())
		s.DifferenceWith(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		s    []int // 入力する値
		t    []int
		want []int // 最終的に求める状態
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{0, 1, 2, 3}, []int{4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6, 123456}, []int{1, 2, 3, 4, 5, 6, 123456}},
		{[]int{0, 1, 2, 3, 123456}, []int{0, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6, 123456}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.s {
			s.Add(value)
		}
		var tSet IntSet
		for _, value := range test.t {
			tSet.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).SymmetricDifference(%v)", s.String(), tSet.String())
		s.SymmetricDifference(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}
