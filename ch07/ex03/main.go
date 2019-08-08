package main

import (
	"crypto/rand"

	"gopl.io/ch4/treesort"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
}
