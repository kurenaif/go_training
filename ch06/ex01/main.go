package main

import (
	"fmt"
	"go_training/ch06/ex01/intset"
)

func main() {
	var x intset.IntSet
	x.Add(1)
	fmt.Println(x.String())
}
