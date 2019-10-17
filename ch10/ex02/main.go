package main

import (
	"go_training/ch10/ex02/unarchive"
	_ "go_training/ch10/ex02/unarchive/tar"
	_ "go_training/ch10/ex02/unarchive/zip"
	"os"
)

type FileFormat struct {
	magic  []byte
	offset int
}

func main() {
	unarchive.ListFormat()
	unarchive.Decode(os.Stdin)
}
