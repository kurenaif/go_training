package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countWords("test.txt")
}

func countWords(path string) {
	counts := make(map[string]int)
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
