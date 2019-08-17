package main

import (
	"fmt"
	"go_training/ch07/ex13/eval"
)

func main() {
	expr, _ := eval.Parse("1+2*3")
	fmt.Println(expr.String())
}
