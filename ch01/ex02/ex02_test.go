package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"one"}, "0 one\n"},
		{[]string{"one", "two", "three"}, "0 one\n1 two\n2 three\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)
		out = new(bytes.Buffer)
		echo(test.args)
		got := out.(*bytes.Buffer).String()
		fmt.Println(got)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
