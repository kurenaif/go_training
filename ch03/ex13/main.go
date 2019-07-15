package main

import (
	"fmt"
)

const (
	KB = 1000      // 1000
	MB = KB * 1000 // 1000000
	GB = MB * 1000 // 1000000000
	TB = GB * 1000 // 1000000000000
	PB = TB * 1000 // 1000000000000000
	EB = PB * 1000 // 1000000000000000000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// fmt.Println(ZB) //overflow!
	// fmt.Println(YB)
}