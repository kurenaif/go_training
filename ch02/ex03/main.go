package main

import (
	"fmt"
	"go_training/ch02/ex03/popcount"
	"go_training/ch02/ex03/popcountloop"
	"strconv"
)

func main() {
	num, _ := strconv.ParseUint("0", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcountloop.PopCount(num))
	num, _ = strconv.ParseUint("0", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
}
