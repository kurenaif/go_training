package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const MOD = 1000000007

var memo []int

func fibo(a int) int {
	if memo[a] != 0 {
		return memo[a]
	}
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	memo[a] = (fibo(a-1) + fibo(a-2)) % MOD
	return memo[a]
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	memo = make([]int, num+1)
	fmt.Println(fibo(num))
}
