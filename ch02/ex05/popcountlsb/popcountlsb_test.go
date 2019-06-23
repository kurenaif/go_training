package popcountlsb

import (
	"fmt"
	"strconv"
	"testing"
)

func myParseUint(str string) uint64 {
	num, _ := strconv.ParseUint(str, 2, 0)
	return num
}

func TestPopCount(t *testing.T) {
	var tests = []struct {
		number uint64
		want   int
	}{
		{myParseUint("0000000000000000000000000000000000000000000000000000000000000000"), 0},
		{myParseUint("1111111111111111111111111111111111111111111111111111111111111111"), 64},
		{myParseUint("1000000000000000000000000000000000000000000000000000000000000001"), 2},
		{myParseUint("1000000010000000000100001000000000000001000000000000000000000001"), 6},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PopCount(%d)", test.number)
		got := PopCount(test.number)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}
