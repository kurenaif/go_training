package main

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"hello", 1},
		{"hello, world", 2},
		{"hello, world\nhello", 3},
		{"", 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(*ByteCounter).Write([]byte(%s))", test.s)
		var c ByteCounter
		c.Write([]byte(test.s))

		if int(c) != test.want {
			t.Errorf("%s = %v want %v", descr, c, test.want)
		}
	}
}
