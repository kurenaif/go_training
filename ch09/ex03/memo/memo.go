// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2019 kurenaif
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import "fmt"

//!+Func

// Func is the type of the function to memoize.
type Func func(key string, done chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
	ok    bool          // 結果が無事に帰ってきたかを保存するためのフラグ(doneは後からcloseできるため、結果が帰った直後の状態を持っておく)
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	done     chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func isClosed(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-get

func getCache(cache map[string]*entry, key string) *entry {
	e := cache[key]
	if e == nil || !e.ok {
		return nil
	}
	return e
}

//!+monitor

// もとのソースコードの意向を組んで、同じサイトへの同時アクセスは1に絞る
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := getCache(cache, req.key)
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e              // ここで代入してるけどe.okをfalseにすると実質cacheしてないことになる
			go e.call(f, req.key, req.done) // call f(key)
		}
		go e.deliver(req, memo)
	}
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, done)
	// Broadcast the ready condition.
	if !isClosed(done) { // 値が帰ったタイミングではcloseされていない=>この値はcacheして良い。
		e.ok = true
	}
	close(e.ready)
}

func (e *entry) deliver(req request, memo *Memo) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	// requestがclosedであれば、cacheされていてもnilを帰す
	if isClosed(req.done) {
		req.response <- result{nil, fmt.Errorf("request canceled")}
		return
	}
	if !e.ok {
		memo.requests <- req // queueに積み直す
		return               // req.responseには返さない
	}
	req.response <- e.res
}

//!-monitor
