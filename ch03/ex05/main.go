package main

import (
	"go_training/ch03/ex05/mandelprot"
	"os"
)

func main() {
	mandelprot.MakeImage(os.Stdout)
}
