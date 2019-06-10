package main

import (
	"os"

	"./lissajous"
)

func main() {
	lissajous.Lissajous(os.Stdout)
}
