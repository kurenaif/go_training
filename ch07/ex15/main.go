package main

import (
	"fmt"
	"go_training/ch07/ex15/eval"
	"strconv"
)

func envInput(vars map[eval.Var]bool) (env eval.Env) {
	env = eval.Env{}

	for v := range vars {
		fmt.Printf("%s: ", v)
		var valueStr string
		fmt.Scanln(&valueStr)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		env[v] = value
		break
	}
	return env
}

func evalInput() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	fmt.Print("Enter expr: ")
	var input string
	fmt.Scanln(&input)
	expr, err := eval.Parse(input)
	if err != nil {
		fmt.Println(err)
	}
	vars := map[eval.Var]bool{}
	expr.Check(vars)
	env := envInput(vars)
	fmt.Println(expr.Eval(env))
}

func main() {
	evalInput()
}
