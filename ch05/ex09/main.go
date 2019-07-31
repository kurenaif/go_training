package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	res := expand("hello world $くれ😣なゐ golang $kijin $seija", func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	})
	fmt.Println(res)
}

func expand(s string, f func(string) string) (res string) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if res != "" {
			res += " "
		}
		txt := scanner.Text()
		if len(txt) > 0 && txt[0] == '$' {
			res += f(txt[1:])
		} else {
			res += txt
		}
	}
	return
}
