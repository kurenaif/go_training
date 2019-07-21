package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func getKinds(r rune) (res []string) {
	if unicode.IsControl(r) {
		res = append(res, "unicode.IsControl")
	}
	if unicode.IsDigit(r) {
		res = append(res, "unicode.IsDigit")
	}
	if unicode.IsGraphic(r) {
		res = append(res, "unicode.IsGraphic")
	}
	if unicode.IsLetter(r) {
		res = append(res, "unicode.IsLetter")
	}
	if unicode.IsLower(r) {
		res = append(res, "unicode.IsLower")
	}
	if unicode.IsMark(r) {
		res = append(res, "unicode.IsMark")
	}
	if unicode.IsNumber(r) {
		res = append(res, "unicode.IsNumber")
	}
	if unicode.IsPrint(r) {
		res = append(res, "unicode.IsPrint")
	}
	if unicode.IsPunct(r) {
		res = append(res, "unicode.IsPunct")
	}
	if unicode.IsSpace(r) {
		res = append(res, "unicode.IsSpace")
	}
	if unicode.IsSymbol(r) {
		res = append(res, "unicode.IsSymbol")
	}
	if unicode.IsTitle(r) {
		res = append(res, "unicode.IsTitle")
	}
	if unicode.IsUpper(r) {
		res = append(res, "unicode.IsUpper")
	}
	return
}

func main() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for _, kind := range getKinds(r) {
			counts[kind]++
		}
		utflen[n]++
	}
	fmt.Printf("property\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
