package main

import (
	"go_training/ch12/ex01/display"
)

type S struct {
	name string
	age  int
}

type T struct {
	name string
	age  int
	arr  [3]int
	s    struct {
		arr2 [2]int
	}
}

func main() {
	a := make(map[[3]int]int)
	k := [3]int{0, 1, 2}
	a[k] = 3
	k = [3]int{1, 1, 2}
	a[k] = 5
	var s S
	s.name = "hello"
	s.age = 32
	b := make(map[S]bool)
	b[s] = true

	c := make(map[T]bool)
	var t T
	t.name = "world"
	t.age = 3
	t.arr = [3]int{1, 2, 3}
	t.s.arr2 = [2]int{1, 2}
	c[t] = true
	display.Display("map", a)
	display.Display("map", b)
	display.Display("map", c)
}
