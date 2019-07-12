package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// 符号付き、小数点付き数字
// フランス式に準拠整数部は,で小数部はスペース区切り
func comma(s string) string {
	var sign byte
	if s[0] == '+' || s[0] == '-' {
		sign = s[0]
		s = s[1:]
	}
	// find period
	integer := s
	fraction := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			integer = s[:i]
			fraction = s[i+1:]
		}
	}

	// ".123 45" みたいなのを防ぐ
	// "0.123 45"みたいにする
	if integer == "" {
		integer = "0"
	}

	// 返り値の生成(signがbyteなため、bufferにした)
	var buf bytes.Buffer
	if sign != 0 {
		buf.WriteByte(sign)
	}
	buf.WriteString(commaInteger(integer))
	if fraction != "" {
		buf.WriteString("." + commaFraction(fraction))
	}
	return buf.String()
}

func commaInteger(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaInteger(s[:n-3]) + "," + s[n-3:]
}

func commaFraction(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + " " + commaFraction(s[3:])
}
