// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 365.

//!+

// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"go_training/ch13/ex04/bzip"
	"io"
	"log"
	"os"
)

func main() {
	w, err := bzip.NewWriter(os.Stdout)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

//!-
