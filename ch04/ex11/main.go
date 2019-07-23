package main

import (
	"go_training/ch04/ex11/create"
	"go_training/ch04/ex11/ghutil"
	"log"
	"os"
)

func main() {
	s, err := ghutil.MessageAndInput("[create]>")
	if err != nil {
		log.Fatal(err)
	}

	if s[:1] == "c" {

		err = create.CreateIssue(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}
