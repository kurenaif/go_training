// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 348.

// Search is a demo of the params.Unpack function.
package main

import (
	"fmt"
	"go_training/ch12/ex11/params"
	"log"
	"net/http"
	"reflect"
)

//!+

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
	res, err := params.Pack(reflect.ValueOf(data), "http://localhost:12345/")
	if err != nil {
		fmt.Fprintf(resp, "Pack error: %v\n", err)
	}
	fmt.Fprintf(resp, "Pack: %v\n", res)
}

//!-

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
+ sleep 1
+ go run main.go
2019/11/19 02:35:38 listen tcp :12345: bind: address already in use
exit status 1
+ curl http://localhost:12345/search
Search: {Labels:[] MaxResults:10 Exact:false}
Pack: http:/localhost:12345/?x=false&&max=10
+ curl 'http://localhost:12345/search?l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:false}
Pack: http:/localhost:12345/?l=golang&l=programming&max=10&x=false
+ curl 'http://localhost:12345/search?l=golang&l=programming&max=100'
Search: {Labels:[golang programming] MaxResults:100 Exact:false}
Pack: http:/localhost:12345/?max=100&x=false&l=golang&l=programming
+ curl 'http://localhost:12345/search?x=true&l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:true}
Pack: http:/localhost:12345/?l=golang&l=programming&max=10&x=true
+ curl 'http://localhost:12345/search?q=hello&x=123'
x: strconv.ParseBool: parsing "123": invalid syntax
+ curl 'http://localhost:12345/search?q=hello&max=lots'
max: strconv.ParseInt: parsing "lots": invalid syntax
*/
