package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		(*c)++
	}
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
