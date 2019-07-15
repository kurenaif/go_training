package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"os"
)

func hash(s string, hashAlgorithm string) (string, error) {
	switch hashAlgorithm {
	case "sha256":
		return fmt.Sprintf("%x", sha256.Sum256([]byte(s))), nil
	case "sha384":
		return fmt.Sprintf("%x", sha512.Sum384([]byte(s))), nil
	case "sha512":
		return fmt.Sprintf("%x", sha512.Sum512([]byte(s))), nil
	}
	return "", errors.New("hashAlgorithm not found")
}

func main() {

	hashAlgorithm := "sha256"
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "sha256":
			hashAlgorithm = os.Args[1]
		case "sha384":
			hashAlgorithm = os.Args[1]
		case "sha512":
			hashAlgorithm = os.Args[1]
		default:
			fmt.Fprintln(os.Stderr, "usage:"+os.Args[0]+"[sha256 or sha384 or sha512]")
			os.Exit(1)
		}

	}

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for input.Scan() {
		txt := input.Text()
		res, err := hash(txt, hashAlgorithm)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(hashAlgorithm, ":", res)
		fmt.Print("> ")
	}
}
