package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestByteBitDiffCount(t *testing.T) {
	var tests = []struct {
		lhsStr string
		rhsStr string
		want   int
	}{
		{"10001001", "10001001", 0},
		{"10001001", "10001000", 1},
		{"00001001", "10001001", 1},
		{"00000000", "11111111", 8},
		{"11111111", "00000000", 8},
		{"00000000", "00000000", 0},
		{"11111111", "11111111", 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("byteBitDiffCount(0b%s, 0b%s)", test.lhsStr, test.rhsStr)
		lhs, _ := strconv.ParseUint(test.lhsStr, 2, 8) // 注意! PraseIntを使用すると、符号bitの都合で最大値が127になる
		rhs, _ := strconv.ParseUint(test.rhsStr, 2, 8)
		got := byteBitDiffCount(byte(lhs), byte(rhs))
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}

func TestHashBitDiffCount(t *testing.T) {
	var tests = []struct {
		lhs  string
		rhs  string
		want int
	}{
		{"abc", "abc", 0},
		{"hello", "world", 112},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("hashBitDiffCount(%s, %s)", test.lhs, test.rhs)
		got := hashBitDiffCount(test.lhs, test.rhs)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}
