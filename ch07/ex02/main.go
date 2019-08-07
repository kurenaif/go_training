package main

import (
	"fmt"
	"io"
	"os"
)

type Counter struct {
	cnt    int64
	writer io.Writer
}

func (c *Counter) Write(p []byte) (int, error) {
	c.cnt += int64(len(p))
	c.writer.Write(p)
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c Counter
	c.writer = w
	return &c, &c.cnt
}

func main() {
	writer, cnt := CountingWriter(os.Stdout)
	writer.Write([]byte("Hello\n"))
	fmt.Println(*cnt)
}
