package main

import (
	"fmt"
	"go_training/ch04/ex10/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1], os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
