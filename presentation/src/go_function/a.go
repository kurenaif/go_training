package main

import "fmt"

func f(a int, b int) int {
	var memo [100]int
	index := 0
	for i := a; i < b; i++ {
		memo[index] = i
		index++
	}

	res := 0
	for ; index > 0; index-- {
		res += memo[index-1]
	}
	return res
}

func main() {
	fmt.Println(f(1, 2))
}
