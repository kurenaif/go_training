package main

import (
	"fmt"
	"go_training/ch01/ex06/lissajous"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := 5

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "query parameter is invalid")
			log.Print(err)
			return
		}
		// Formの型: type Values map[string][]string
		if vals, isExist := r.Form["cycles"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"cylcles\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			s, err := strconv.Atoi(val)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			cycles = s
		}
		lissajous.Lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
