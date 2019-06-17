// コマンドライン引数が0のとき、標準入力から受けとる
// 標準入力で入力する際は、スペース区切りでも入力ができる
// Ctrl+Cで終了する
// コマンドライン引数が1以上のとき、与えられた引数に対応した出力を吐いて終了する
package main

import (
	"bufio"
	"fmt"
	"go_training/ch02/ex01/tempconv"
	"go_training/ch02/ex02/lengthconv"
	"go_training/ch02/ex02/weightconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			input := stdin.Text()
			printConv(strings.Split(input, " "))
		}
	} else {
		printConv(args)
	}
}

func printConv(args []string) {
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch02:ex02: %v\n", err)
			os.Exit(1)
		}
		celsius := tempconv.Celsius(t)
		fahrenheit := tempconv.Fahrenheit(t)
		meter := lengthconv.Meter(t)
		feet := lengthconv.Feet(t)
		kilogram := weightconv.Kilogram(t)
		pounds := weightconv.Pounds(t)
		fmt.Printf("%s\t=\t%s,\t%s\t=\t%s\n", fahrenheit, tempconv.FToC(fahrenheit), celsius, tempconv.CToF(celsius))
		fmt.Printf("%s\t=\t%s,\t%s\t=\t%s\n", meter, lengthconv.MToF(meter), meter, lengthconv.FToM(feet))
		fmt.Printf("%s\t=\t%s,\t%s\t=\t%s\n", kilogram, weightconv.KToP(kilogram), pounds, weightconv.PToK(pounds))
	}
}
