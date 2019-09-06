package main

import "fmt"

func double(x int) (res int) {
	defer func() {
		if p := recover(); p != nil {
			res = p.(int)
		}
	}()
	panic(x * x)
}

func main() {
	fmt.Println(double(3))
}
