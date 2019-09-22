// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"fmt"
	"go_training/ch02/ex03/popcount"
	"strconv"
	"testing"
)

// -- Alternative implementations --

// 1のケース後半の演習の比較用
func TestPopCount(t *testing.T) {
	type Test struct {
		numStr string
		want   int
	}
	tests := []Test{
		{"1111111111", 10},
		{"1111111111111111111111111111111111111111111111111111111111111111", 64},
		{"0", 0},
	}

	done := make(chan struct{}, len(tests))

	for _, test := range tests {
		go func(test Test) {
			descr := fmt.Sprintf("popcount.PopCount(%s)", test.numStr)
			num, _ := strconv.ParseUint(test.numStr, 2, 0)
			got := popcount.PopCount(num)
			fmt.Printf("%s = %d\n", descr, got)
			if got != test.want {
				t.Errorf("%s = %d, want %d", descr, got, test.want)
			}
			done <- struct{}{}
		}(test)
	}

	for i := 0; i < len(tests); i++ {
		<-done
	}
}
