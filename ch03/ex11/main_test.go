package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567890", "1,234,567,890"},
		{"234567890", "234,567,890"},
		{"34567890", "34,567,890"},
		{"34567890.", "34,567,890"},
		{"34567890.1", "34,567,890.1"},
		{"34567890.12", "34,567,890.12"},
		{"34567890.123", "34,567,890.123"},
		{"34567890.1234", "34,567,890.123 4"},
		{"34567890.1234567890", "34,567,890.123 456 789 0"},
		{"+34567890.1234567890", "+34,567,890.123 456 789 0"},
		{"-34567890.1234567890", "-34,567,890.123 456 789 0"},
		{".1234567890", "0.123 456 789 0"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("comma(%q)", test.s)
		got := comma(test.s)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
