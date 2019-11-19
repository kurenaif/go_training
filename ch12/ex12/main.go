// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 348.

// Search is a demo of the params.Unpack function.
package main

import (
	"fmt"
	"go_training/ch12/ex12/params"
	"log"
	"net/http"
	"strings"
)

//!+

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l" validator:"LavelValidate"`
		Str        string   `http:"s" validator:"StrValidate"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	unpacker := params.NewUnpacker()

	unpacker.RegistValidator("StrValidate", func(s interface{}) error {
		str, ok := s.(string)
		if !ok {
			return fmt.Errorf("StrValidate given not string")
		}
		if !strings.HasPrefix(str, "hello") {
			return fmt.Errorf("%q is not start from \"hello\"", str)
		}
		return nil
	})

	unpacker.RegistValidator("LavelValidate", func(s interface{}) error {
		strs, ok := s.([]string)
		if !ok {
			return fmt.Errorf("StrValidate given not []string")
		}
		if len(strs) != 2 {
			return fmt.Errorf("len of Label(%q) is not 2", strs)
		}
		return nil
	})

	if err := unpacker.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

//!-

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
$ curl 'http://localhost:12345/search?s=hello&x=true&l=golang&l=prog'
Search: {Labels:[golang prog] Str:hello MaxResults:10 Exact:true}
$ 'http://localhost:12345/search?s=hello&x=true&l=golang'
name validator error: len of Label(["golang"]) is not 2
$ 'http://localhost:12345/search?s=hell&x=true&l=golang&l=prog'
name validator error: "hell" is not start from "hello"
*/
