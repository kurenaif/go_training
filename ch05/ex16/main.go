package main

import "fmt"

func join(sep string, strs ...string) (res string) {
	s := ""
	for _, str := range strs {
		res += s + str
		s = sep
	}
	return res
}

func main() {
	fmt.Println(join(",", "a", "b", "c"))
}
