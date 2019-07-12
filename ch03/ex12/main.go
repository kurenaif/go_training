package main

import (
	"fmt"
	"os"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("  %s\n", comma(os.Args[i]))
	// }
	if len(os.Args) < 3 {
		fmt.Println("usage: " + os.Args[0] + " str1 str2")
		os.Exit(1)
	}
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(lhs string, rhs string) bool {
	leftRunes := runeCount(lhs)
	rightRunes := runeCount(rhs)

	for k, v := range leftRunes {
		val := rightRunes[k]
		rightRunes[k] = val - v
	}

	for _, v := range rightRunes {
		if v != 0 {
			return false
		}
	}
	return true
}

func runeCount(s string) map[rune]int {
	cnt := map[rune]int{}
	for _, r := range s {
		val := cnt[r]
		cnt[r] = val + 1
	}
	return cnt
}
