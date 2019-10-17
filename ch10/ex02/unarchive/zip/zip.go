package zip

import (
	"fmt"
	"go_training/ch10/ex02/unarchive"
	"io"
	"os"
)

func decode(io.Reader) (os.FileInfo, error) {
	fmt.Println("zip decoded!")
	return nil, nil
}

func init() {
	unarchive.RegisterFormat("zip", []byte{0x50, 0x4B}, 0x0, decode)
}
