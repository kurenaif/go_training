package main

import (
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func EchoFor(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func EchoJoin(args []string) {
	strings.Join(args, " ")
	// fmt.Fprintln(out, s)
}

func main() {
	EchoFor([]string{"hello"})
	EchoJoin([]string{"hello"})
}
