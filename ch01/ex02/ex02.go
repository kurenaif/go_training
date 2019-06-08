package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func echo(args []string) {
	for index, arg := range args {
		fmt.Fprintln(out, index, arg)
	}
}

func main() {
	echo(os.Args[0:])
}
