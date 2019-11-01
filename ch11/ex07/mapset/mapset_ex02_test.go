package mapset

import (
	"fmt"
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		values    []int // 入力する値
		addValues []int // 削除する値
		want      []int // 最終的に求める状態
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{}, []int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{0, 1, 2}, []int{1, 2, 3}, []int{0, 1, 2, 3}},
	}

	for _, test := range tests {
		var s IntSet
		for _, value := range test.values {
			s.Add(value)
		}

		descr := fmt.Sprintf("(*IntSet(%v)).AddAll(%v)", s.String(), test.addValues)
		s.AddAll(test.addValues...)

		var want IntSet
		for _, value := range test.want {
			want.Add(value)
		}

		if !isSame(s, want) {
			t.Errorf("%s = %v want %v", descr, s.String(), want.String())
		}
	}
}
