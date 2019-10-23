package main

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {

	var tests = []struct {
		str       string
		count     map[rune]int
		utfLenMap map[int]int
		invalid   int
	}{
		{"", map[rune]int{}, map[int]int{}, 0},
		{"AAABB", map[rune]int{'A': 3, 'B': 2}, map[int]int{1: 5}, 0},
		{"あああいいいううう", map[rune]int{'あ': 3, 'い': 3, 'う': 3}, map[int]int{3: 9}, 0},
		{"Aあ😃", map[rune]int{'A': 1, 'あ': 1, '😃': 1}, map[int]int{1: 1, 3: 1, 4: 1}, 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("count, urlLen, invalid := CharCount(%q)", test.str)
		t.Logf(" \n testing... \n      %s \n", descr)
		gotCounts, gotUtfLen, gotInvalid := CharCount(bufio.NewReader(strings.NewReader(test.str)))
		utfLen := [utf8.UTFMax + 1]int{}
		for k, v := range test.utfLenMap {
			utfLen[k] = v
		}

		if !reflect.DeepEqual(gotCounts, test.count) {
			t.Errorf("count = %x, want %x", gotCounts, test.count)
		}

		if !reflect.DeepEqual(gotUtfLen, utfLen) {
			t.Errorf("urlLen = %x, want %x", gotUtfLen, utfLen)
		}

		if gotInvalid != test.invalid {
			t.Errorf("invalid = %x, want %x", gotInvalid, test.invalid)
		}
	}
}

func TestInvalidChar(t *testing.T) {

	// strは1byte 以上あることを想定
	var tests = []struct {
		str       string
		count     map[rune]int
		utfLenMap map[int]int
		invalid   int
	}{
		{"あああ", map[rune]int{0: 1, 'あ': 2}, map[int]int{1: 1, 3: 2}, 2},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("count, urlLen, invalid := CharCount(%q)", test.str)
		vec := []byte(test.str)
		vec[0] = 0
		t.Logf(" \n testing... \n      %s \n", descr)
		gotCounts, gotUtfLen, gotInvalid := CharCount(bufio.NewReader(bytes.NewReader(vec)))
		utfLen := [utf8.UTFMax + 1]int{}
		for k, v := range test.utfLenMap {
			utfLen[k] = v
		}

		if !reflect.DeepEqual(gotCounts, test.count) {
			t.Errorf("count = %x, want %x", gotCounts, test.count)
		}

		if !reflect.DeepEqual(gotUtfLen, utfLen) {
			t.Errorf("urlLen = %x, want %x", gotUtfLen, utfLen)
		}

		if gotInvalid != test.invalid {
			t.Errorf("invalid = %x, want %x", gotInvalid, test.invalid)
		}
	}
}
