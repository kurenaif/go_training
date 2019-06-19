package main

import (
	"fmt"
	"go_training/ch02/ex03/popcount"
	"strconv"
)

func main() {
	num, _ := strconv.ParseUint("010101", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
	num, _ = strconv.ParseUint("111111", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
	num, _ = strconv.ParseUint("0", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
	num, _ = strconv.ParseUint("000001", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
}
