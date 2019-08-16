// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	fmt.Println("http://localhost:8000/list")
	fmt.Println("http://localhost:8000/price?item=XXXX")
	fmt.Println("http://localhost:8000/create?item=XXXX&price=YYYY")
	fmt.Println("http://localhost:8000/update?item=XXXX&price=YYYY")
	fmt.Println("http://localhost:8000/delete?item=XXXX")
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.deleteItem)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

var itemValueTable = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $item, $value := .}}
<tr>
  <td>{{$item}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var res bytes.Buffer
	itemValueTable.Execute(&res, db)
	fmt.Fprint(w, res.String())
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	if item == "" {
		fmt.Fprintf(w, "field 'item' must be setted.")
		return
	}
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		fmt.Fprintf(w, "field 'price' error: %s", err)
		return
	}
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%s is already exist.", item)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s setted\n", item, dollars(price))
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	if item == "" {
		fmt.Fprintf(w, "field 'item' must be setted.")
		return
	}
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		fmt.Fprintf(w, "field 'price' error: %s", err)
		return
	}
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s is not exist.", item)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s setted\n", item, dollars(price))
}

func (db database) deleteItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s is not exist.", item)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	delete(db, item)
	fmt.Fprintf(w, "%s is deleted", item)
}
