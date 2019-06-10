package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// 応答がないウェブサイトがある場合: 結果を受け取ることができず、 fmt.Fprintln(writer, <-ch)で処理が止まる。
// クライアント側でもtimeoutを実装してゴルーチンにcancelさせる手法を導入する必要があると考えられる。

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Println("usage: ex11.go urllist_filename")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stdout, "ex11: %v\n", err)
		os.Exit(1)
	}
	urlList := strings.Split(string(data), "\n")
	start := time.Now()
	fetchAll(os.Stdout, urlList)
	fmt.Fprintf(os.Stdout, "%.2fs elapced\n", time.Since(start).Seconds())
}

func fetchAll(writer io.Writer, urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range urls {
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
