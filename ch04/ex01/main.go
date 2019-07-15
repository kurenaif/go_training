package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func byteBitDiffCount(lhs byte, rhs byte) (cnt int) {
	for i := byte(0); i < 8; i++ {
		if (lhs & 1) != (rhs & 1) {
			cnt++
		}
		lhs >>= 1
		rhs >>= 1
	}
	return
}

func hashBitDiffCount(s1 string, s2 string) (cnt int) {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	for i := 0; i < len(c1); i++ {
		cnt += byteBitDiffCount(c1[i], c2[i])
	}
	return
}

func main() {
	fmt.Println(hashBitDiffCount(os.Args[1], os.Args[2]))
}
