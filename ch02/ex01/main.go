package main

import (
	"fmt"
	"go_training/ch02/ex01/tempconv"
)

func main() {
	fmt.Println("Hello world")
	c := tempconv.Celsius(3.2)
	f := tempconv.Fahrenheit(4.2)
	k := tempconv.Kelvin(5.3)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(k)
	fmt.Println(tempconv.KToC(tempconv.CToK(c) + tempconv.FToK(f) + k))
}
