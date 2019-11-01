package intset

import (
	"fmt"
	"testing"
)

// 高速版Addの確認 Addと同じであれば良し
func TestFastAdd(t *testing.T) {
	tests := []struct {
		values []int // 入力する値
	}{
		{[]int{0, 1, 2, 3, 4}},
		{[]int{0, 0, 0, 0, 0}},
		{[]int{}},
		{[]int{0, 128, 256, 123456}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IntSet.FastAdd(%d)", test.values)
		var s IntSet
		var s2 IntSet

		for _, value := range test.values {
			s.Add(value)
			s2.FastAdd(value)
		}

		if !isSame(s, s2) {
			t.Errorf("%s = %s, got %s\n", descr, s2.String(), s.String())
		}
	}
}

func TestFastUnionWith(t *testing.T) {
	tests := []struct {
		s []int // 入力する値
		t []int
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{0, 1, 2, 3}, []int{4, 5, 6}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6}},
		{[]int{0, 1, 2, 3}, []int{0, 4, 5, 6, 123456}},
		{[]int{0, 1, 2, 3, 123456}, []int{0, 4, 5, 6}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IntSet(%d).FastUnionWith(%d)", test.s, test.t)

		var s IntSet
		var s2 IntSet
		for _, value := range test.s {
			s.Add(value)
			s2.Add(value)
		}
		var tSet IntSet
		var tSet2 IntSet
		for _, value := range test.t {
			tSet.Add(value)
			tSet2.Add(value)
		}
		s.FastUnionWith(&tSet)
		s2.UnionWith(&tSet2)

		if !isSame(s, s2) {
			t.Errorf("%s = %s want %s", descr, s.String(), s2.String())
		}
	}

}

func TestFastRemove(t *testing.T) {
	tests := []struct {
		values      []int // 入力する値
		removeValue int   // 削除する値
		want        []int // 最終的に求める状態
	}{
		{[]int{0, 1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
		{[]int{0}, 0, []int{}},
		{[]int{0}, 1, []int{0}},
		{[]int{0}, 123456, []int{0}},
		{[]int{}, 0, []int{}},
		{[]int{0, 128, 256, 123456}, 128, []int{0, 256, 123456}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.values {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).Remove(%d)", s.String(), test.removeValue)
		s.FastRemove(test.removeValue)
		got := s

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(got, want) {
			t.Errorf("%s = %v want %v", descr, got.String(), want.String())
		}
	}
}

func TestFastIntersectWith(t *testing.T) {
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
		s.FastIntersectWith(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}

func TestFastDifferenceWith(t *testing.T) {
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
		s.FastDifferenceWith(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}

func TestFastSymmetricDifference(t *testing.T) {
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
		s.FastSymmetricDifference(&tSet)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %s want %s", descr, s.String(), want.String())
		}
	}

}
