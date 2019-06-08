package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"one"}, "one\n"},
		{[]string{"one", "two", "three"}, "one two three\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)
		out = new(bytes.Buffer)
		args := append(os.Args[0:1], test.args...)
		want := os.Args[0] + " " + test.want
		echo(args)
		got := out.(*bytes.Buffer).String()
		if got != want {
			t.Errorf("%s = %q, want %q", descr, got, want)
		}
	}
}
