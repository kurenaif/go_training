package main

import "fmt"

func unique(ss []string) []string {
	i := 0
	for _, s := range ss {
		if i == 0 || ss[i-1] != s {
			ss[i] = s
			i++
		}
	}
	return ss[:i]
}

func main() {
	fmt.Println(unique([]string{"hello", "hello", "hello", "world", "world", "hello", "world"}))
}
