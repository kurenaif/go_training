package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var out io.Writer = os.Stdout

func main() {
	outputCountLines(os.Args[1:])
}

func outputCountLines(files []string) {
	counts := make(map[string]map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// mapを直に使用すると、hashTableのため、順番がバラバラになってテストがつらい
	// 出力はkeyの昇順で固定 O(N|S|log(N))
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 昇順にfilenameごとに出力
	// lineの総和が2以上で出力
	for _, line := range keys {
		filenameCount := counts[line]
		countSum := 0
		message := ""
		for filename, count := range filenameCount {
			countSum += count
			message += fmt.Sprintf("%s\t%s\t%d\n", line, filename, count)
		}
		if countSum > 1 {
			fmt.Fprintln(out, line+"(sum:", countSum, "):")
			fmt.Fprint(out, message)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}
