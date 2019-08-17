package main

import (
	"fmt"
	"go_training/ch07/ex14/eval"
)

func main() {
	expr, err := eval.Parse("~(x ,4)~")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(expr.Eval(eval.Env{"x": 3}))
}
