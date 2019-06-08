package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func echo(args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}

func main() {
	echo(os.Args[0:])
}
