// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done chan struct{}) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M) {
	//!+seq
	memo := map[string]bool{}
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		memo[url] = true
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		done := make(chan struct{})
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func SequentialCloseFirstCase(t *testing.T, m M) {
	memo := map[string]bool{}
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		if ok := memo[url]; !ok {
			close(done)
		}
		memo[url] = true
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func SequentialCloseSecondCase(t *testing.T, m M) {
	memo := map[string]int{}
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		cnt := memo[url]
		if cnt == 1 {
			close(done)
		}
		memo[url] = cnt + 1
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func SequentialCloseThirdCase(t *testing.T, m M) {
	memo := map[string]int{}
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		cnt := memo[url]
		if cnt == 2 {
			close(done)
		}
		memo[url] = cnt + 1
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

/*
=== RUN   Test
https://golang.org, 392.509034ms, 11060 bytes
https://godoc.org, 315.038778ms, 6841 bytes
https://play.golang.org, 577.059456ms, 6011 bytes
http://gopl.io, 279.403735ms, 4154 bytes
https://golang.org, 287.606µs, 11060 bytes
https://godoc.org, 255.361µs, 6841 bytes
https://play.golang.org, 251.876µs, 6011 bytes
http://gopl.io, 253.143µs, 4154 bytes
https://golang.org, 245.925µs, 11060 bytes
https://godoc.org, 194.272µs, 6841 bytes
https://play.golang.org, 86.184µs, 6011 bytes
http://gopl.io, 85.203µs, 4154 bytes
--- PASS: Test (1.57s)
=== RUN   TestConcurrent
http://gopl.io, 139.608185ms, 4154 bytes
https://play.golang.org, 139.060263ms, 6011 bytes
https://golang.org, 141.171037ms, 11060 bytes
https://golang.org, 142.227232ms, 11060 bytes
https://golang.org, 141.250326ms, 11060 bytes
https://godoc.org, 199.86181ms, 6841 bytes
https://godoc.org, 207.454157ms, 6841 bytes
http://gopl.io, 263.460818ms, 4154 bytes
http://gopl.io, 270.301119ms, 4154 bytes
https://play.golang.org, 500.117068ms, 6011 bytes
https://play.golang.org, 507.754899ms, 6011 bytes
https://godoc.org, 655.830018ms, 6841 bytes
--- PASS: TestConcurrent (0.66s)
=== RUN   TestCloseFirst
2019/09/25 01:12:55 request canceled
2019/09/25 01:12:55 request canceled
2019/09/25 01:12:55 request canceled
2019/09/25 01:12:55 request canceled
https://golang.org, 145.072569ms, 11060 bytes
https://godoc.org, 201.516588ms, 6841 bytes
https://play.golang.org, 498.390874ms, 6011 bytes
http://gopl.io, 149.389319ms, 4154 bytes
https://golang.org, 118.822µs, 11060 bytes
https://godoc.org, 112.441µs, 6841 bytes
https://play.golang.org, 121.117µs, 6011 bytes
http://gopl.io, 119.082µs, 4154 bytes
--- PASS: TestCloseFirst (1.00s)
=== RUN   TestCloseSecond
https://golang.org, 145.359063ms, 11060 bytes
https://godoc.org, 204.732211ms, 6841 bytes
https://play.golang.org, 498.596478ms, 6011 bytes
http://gopl.io, 146.330127ms, 4154 bytes
2019/09/25 01:12:57 request canceled
2019/09/25 01:12:57 request canceled
2019/09/25 01:12:57 request canceled
2019/09/25 01:12:57 request canceled
https://golang.org, 59.184µs, 11060 bytes
https://godoc.org, 66.598µs, 6841 bytes
https://play.golang.org, 57.18µs, 6011 bytes
http://gopl.io, 77.475µs, 4154 bytes
--- PASS: TestCloseSecond (1.00s)
=== RUN   TestCloseThird
https://golang.org, 141.662757ms, 11060 bytes
https://godoc.org, 215.344755ms, 6841 bytes
https://play.golang.org, 524.518569ms, 6011 bytes
http://gopl.io, 145.3952ms, 4154 bytes
https://golang.org, 67.433µs, 11060 bytes
https://godoc.org, 68.545µs, 6841 bytes
https://play.golang.org, 80.551µs, 6011 bytes
http://gopl.io, 97.377µs, 4154 bytes
2019/09/25 01:12:58 request canceled
2019/09/25 01:12:58 request canceled
2019/09/25 01:12:58 request canceled
2019/09/25 01:12:58 request canceled
--- PASS: TestCloseThird (1.03s)
PASS
ok  	go_training/ch09/ex03/memo	6.262s
*/
