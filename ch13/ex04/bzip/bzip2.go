// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 362.
//
// The version of this program that appeared in the first and second
// printings did not comply with the proposed rules for passing
// pointers between Go and C, described here:
// https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
//
// The rules forbid a C function like bz2compress from storing 'in'
// and 'out' (pointers to variables allocated by Go) into the Go
// variable 's', even temporarily.
//
// The version below, which appears in the third printing, has been
// corrected.  To comply with the rules, the bz_stream variable must
// be allocated by C code.  We have introduced two C functions,
// bz2alloc and bz2free, to allocate and free instances of the
// bz_stream type.  Also, we have changed bz2compress so that before
// it returns, it clears the fields of the bz_stream that contain
// pointers to Go variables.

//!+

// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
#include <stdlib.h>
bz_stream* bz2alloc() { return calloc(1, sizeof(bz_stream)); }
int bz2compress(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen);
void bz2free(bz_stream* s) { free(s); }
*/

import (
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type writer struct {
	w    io.Writer
	file *os.File
}

var mu sync.Mutex

const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var w writer
	file, err := os.Create(filepath.Join("/tmp", randString(100)))
	if err != nil {
		return nil, err
	}
	w.file = file
	w.w = out
	return &w, nil
}

//!-

//!+write
func (w *writer) Write(data []byte) (int, error) {
	return w.file.Write(data)
}

//!-write

//!+close
// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()
	err := exec.Command("/bin/bzip2", w.file.Name()).Run()
	if err != nil {
		return err
	}

	out, err := os.Open(w.file.Name() + ".bz2")
	if err != nil {
		return err
	}

	io.Copy(w.w, out)
	return nil
}

//!-close
