package main

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	tests := []struct {
		str  string
		size int
	}{
		{"abc鬼人正邪😣abc", 1},
		{"abc鬼人正邪😣abc", 1000},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("NewReader(%s)", test.str)
		reader := NewReader(test.str)
		readLength := 0
		bts := []byte{}
		for readLength < len(test.str) {
			buffer := make([]byte, test.size)
			size, _ := reader.Read(buffer)
			readLength += size
			bts = append(bts, buffer[:size]...)
		}
		got := string(bts)
		if got != test.str {
			t.Errorf("testing %s, got = %s, want = %s", descr, got, test.str)
		}
	}
}
