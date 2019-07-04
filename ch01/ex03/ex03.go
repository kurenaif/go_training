package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func EchoFor(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func EchoJoin(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(EchoFor([]string{"hello", "world"}))
	fmt.Println(EchoJoin([]string{"hello", "world"}))
}
