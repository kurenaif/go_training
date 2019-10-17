package tar

import (
	"fmt"
	"go_training/ch10/ex02/unarchive"
	"io"
	"os"
)

func decode(io.Reader) (os.FileInfo, error) {
	fmt.Println("tar decoded!")
	return nil, nil
}

func init() {
	unarchive.RegisterFormat("tar", []byte{0x75, 0x73, 0x74, 0x61, 0x72}, 0x101, decode)
}
