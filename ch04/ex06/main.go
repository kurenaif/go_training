package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func compressSpace(bs []byte) []byte {
	i := 0
	for j := 0; j < len(bs); {
		r, size := utf8.DecodeRune(bs[j:])
		last, _ := utf8.DecodeLastRune(bs[:i])             // 最後のruneを取得
		if !unicode.IsSpace(last) || !unicode.IsSpace(r) { // !(prev == space and cur == space) 文字をskipしない条件
			if unicode.IsSpace(r) { // 空白文字を" "に圧縮する
				r = ' '
			}
			_ = utf8.EncodeRune(bs[i:], r) // => size i += size
			i += size
		}
		j += size
	}
	return bs[:i]
}

func main() {
	rs := []rune{' ', '\t', 'H', 'e', 'l', 'l', 'o', '\t', ' ', '　', '\t', '世', '界', '\r', '\v'}
	bs := []byte(string(rs))
	fmt.Println(string(compressSpace(bs)))
}
