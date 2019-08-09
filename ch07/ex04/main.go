package main

import (
	"fmt"
	"io"
)

type MyBytes struct {
	seek  int
	bytes []byte
}

func (mb *MyBytes) Read(p []byte) (int, error) {
	n := copy(p, mb.bytes[mb.seek:])
	mb.seek += n
	return n, nil
}

func NewReader(s string) io.Reader {
	mb := MyBytes{0, []byte(s)}
	return &mb
}

func main() {
	str := "Hello World"

	reader := NewReader(str)
	readLength := 0
	bts := []byte{}
	for readLength < len(str) {
		buffer := make([]byte, 5)
		size, _ := reader.Read(buffer)
		readLength += size
		fmt.Println(buffer)
		bts = append(bts, buffer[:size]...)
	}
	fmt.Println(string(bts))
}
