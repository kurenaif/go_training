package mapset

import (
	"fmt"
	"testing"
)

// Stringが常に昇順であることを仮定した比較方法
func isSameSlice(lhs []uint, rhs []uint) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, _ := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func TestElems(t *testing.T) {
	tests := []struct {
		s    []int  // 入力する値
		want []uint // 入力する値
	}{
		{[]int{0, 1, 2, 3, 4}, []uint{0, 1, 2, 3, 4}},
		{[]int{}, []uint{}},
		{[]int{0, 0, 0, 1, 2}, []uint{0, 1, 2}},
		{[]int{5, 1, 3, 2, 4}, []uint{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.s {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).Elemns()", s.String())

		elems := s.Elems()

		if !isSameSlice(test.want, elems) {
			t.Errorf("%s = %v want %v", descr, elems, test.want)
		}
	}

}
