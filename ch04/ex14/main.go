// note: https://stackoverflow.com/questions/41176355/go-template-name

package main

import (
	"fmt"
	"go_training/ch04/ex14/github"
	"log"
	"net/http"
)

func main() {
	cache := make(map[string]string)
	fmt.Println("http://localhost:8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<a href=http://localhost:8000/issues> issue </a> <br>
		<a href=http://localhost:8000/milestones> milestones </a> <br>
		<a href=http://localhost:8000/users> users </a>`)
	})
	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		page, ok := cache["issues"]
		if !ok {
			p, err := github.GetIssuePage()
			if err != nil {
				log.Fatal(err)
			}
			page = p
			cache["issues"] = page
		}
		fmt.Fprintf(w, page)
	})
	http.HandleFunc("/milestones", func(w http.ResponseWriter, r *http.Request) {
		page, ok := cache["milestones"]
		if !ok {
			p, err := github.GetMilestonePage()
			if err != nil {
				log.Fatal(err)
			}
			page = p
			cache["milestones"] = page
		}
		fmt.Fprintf(w, page)
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		page, ok := cache["users"]
		if !ok {
			p, err := github.GetMilestonePage()
			if err != nil {
				log.Fatal(err)
			}
			page = p
			cache["users"] = page
		}
		fmt.Fprintf(w, page)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
