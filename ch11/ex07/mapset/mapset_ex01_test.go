package mapset

import (
	"fmt"
	"testing"
)

func isSame(lhs IntSet, rhs IntSet) bool {
	return lhs.String() == rhs.String() // DeepEqualだとnilが…
}

func TestLen(t *testing.T) {
	tests := []struct {
		values []int // 入力する値
		want   int   // これのすべてがtrueでなければならない
	}{
		{[]int{0, 1, 2, 3, 4}, 5},
		{[]int{0, 0, 0, 0, 0}, 1},
		{[]int{}, 0},
		{[]int{0, 128, 256, 123456}, 4},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.values {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).Len()", s.String())
		got := s.Len()
		if got != test.want {
			t.Errorf("%s = %d want %d", descr, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
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
		s.Remove(test.removeValue)
		got := s

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		// fmt.Println(reflect.DeepEqual(got.set, want.set))
		if !isSame(got, want) {
			t.Errorf("%s = %v want %v", descr, got.String(), want.String())
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		values []int // 入力する値
	}{
		{[]int{0, 1, 2, 3, 4}},
		{[]int{0}},
		{[]int{}},
		{[]int{0, 128, 256, 123456}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.values {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).Clear()", s.String())
		s.Clear()
		got := s

		var want IntSet

		if !isSame(got, want) {
			t.Errorf("%s = %v want %v", descr, got.String(), want.String())
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		values []int // 入力する値
	}{
		{[]int{0, 1, 2, 3, 4}},
		{[]int{0}},
		{[]int{}},
		{[]int{0, 128, 256, 123456}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.values {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).Clear()", s.String())

		got := s.Copy()

		// 中身がcopyできていない
		if !isSame(*got, s) {
			t.Errorf("%s = %v want %v", descr, got.String(), (*got).String())
		}

		if got == &s {
			t.Errorf("source and destination IntSets are the same address.")
		}

		// 試しに変な値をAdd
		got.Add(1515151515)
		// 片方だけAddしたのにSameなのはおかしい
		if isSame(*got, s) {
			t.Errorf("%s = %q want %q", descr, got.String(), (*got).String())
		}
	}
}
