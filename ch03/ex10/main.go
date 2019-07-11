package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	size := n % 3
	if size == 0 {
		size = 3
	}

	var buf bytes.Buffer
	index := 0
	for index < n {
		fmt.Fprintf(&buf, "%s", s[index:index+size])
		index += size
		size = 3

		// 最後の,は表示しない
		if index < n {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}
