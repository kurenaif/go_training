package main

import (
	"go_training/ch04/ex11/close"
	"go_training/ch04/ex11/create"
	"go_training/ch04/ex11/ghutil"
	"go_training/ch04/ex11/list"
	"go_training/ch04/ex11/patch"
	"go_training/ch04/ex11/read"
	"log"
	"os"
	"strconv"
)

func main() {
	s, err := ghutil.MessageAndInput("[create, list, read, patch, close]>")
	if err != nil {
		log.Fatal(err)
	}

	token := ""
	if len(os.Args) > 1 {
		token = os.Args[1]
	} else {
		s, err := ghutil.MessageAndInput("token > ")
		if err != nil {
			log.Fatal(err)
		}
		token = s
	}

	if s[:2] == "create" {
		err = create.CreateIssue(token)
		if err != nil {
			log.Fatal(err)
		}
	} else if s[:1] == "l" {
		err = list.ListIssue(token)
		if err != nil {
			log.Fatal(err)
		}
	} else if s[:1] == "r" {
		numberStr, err := ghutil.MessageAndInput("issue number > ")
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		err = read.ReadIssue(token, number)
		if err != nil {
			log.Fatal(err)
		}
	} else if s[:1] == "p" {
		numberStr, err := ghutil.MessageAndInput("issue number > ")
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		err = patch.PatchIssue(token, number)
		if err != nil {
			log.Fatal(err)
		}
	} else if s[:2] == "cl" {
		numberStr, err := ghutil.MessageAndInput("issue number > ")
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		err = close.CloseIssue(token, number)
		if err != nil {
			log.Fatal(err)
		}
	}
}
