package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	file, err := os.Create("out1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex10: %v\n", err)
		os.Exit(1)
	}
	start := time.Now()
	fetchAll(file, os.Args[1:])
	fmt.Fprintf(file, "%.2fs elapced\n", time.Since(start).Seconds())

	file, err = os.Create("out2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex10: %v\n", err)
		os.Exit(1)
	}
	start = time.Now()
	fetchAll(file, os.Args[1:])
	fmt.Fprintf(file, "%.2fs elapced\n", time.Since(start).Seconds())
}

func fetchAll(writer io.Writer, urls []string) {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(writer, <-ch)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}
