package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var done = make(chan struct{})

func mirroredQuery(urls []string) string {
	if len(urls) == 0 { //引数のlenが0のときはchanelが0になり、永遠に受信を待ち続けることに成る
		return ""
	}
	reqs := []*http.Request{}
	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		req.Cancel = done
		reqs = append(reqs, req)
	}

	responses := make(chan string, len(reqs))
	for _, req := range reqs {
		go func(req *http.Request) {
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Print(err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			responses <- string(bodyBytes)
		}(req)
	}
	body := <-responses
	close(done)
	return body
}

func main() {
	fmt.Println(mirroredQuery(os.Args[1:]))
}
