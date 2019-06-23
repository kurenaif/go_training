package main

import (
	"fmt"
	"go_training/ch02/ex03/popcount"
	"go_training/ch02/ex03/popcountloop"
	"go_training/ch02/ex04/popcountbitshift"
	"go_training/ch02/ex05/popcountlsb"
	"strconv"
)

func main() {
	num, _ := strconv.ParseUint("0101010", 2, 0)
	fmt.Println(strconv.FormatUint(num, 2), ":", popcountloop.PopCount(num))
	fmt.Println(strconv.FormatUint(num, 2), ":", popcount.PopCount(num))
	fmt.Println(strconv.FormatUint(num, 2), ":", popcountbitshift.PopCount(num))
	fmt.Println(strconv.FormatUint(num, 2), ":", popcountlsb.PopCount(num))
}
