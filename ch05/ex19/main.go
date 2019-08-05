package main

import "fmt"

func double(x int) {
	panic(x * x)
}

func main() {
	defer func() {
		p := recover()
		fmt.Println(p)
	}()

	double(3)
}
