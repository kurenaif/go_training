package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileSizes struct {
	rootDir string
	size    int64
}

type FilesBytes struct {
	nfiles int64
	nbytes int64
}

var vFlag = flag.Bool("v", false, "show verbose progress messages")

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan FileSizes)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(50 * time.Millisecond)
	}

	rootFB := map[string]FilesBytes{}
loop:
	for {
		select {
		case fileSize, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			fb := rootFB[fileSize.rootDir]
			fb.nfiles++
			fb.nbytes += fileSize.size
			rootFB[fileSize.rootDir] = fb
		case <-tick:
			printDiskUsage(rootFB)
		}
	}

	fmt.Println("--------------------result--------------------")
	printDiskUsage(rootFB) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(rootFB map[string]FilesBytes) {
	for root, fb := range rootFB {
		fmt.Printf("%s: %d files  %.1f GB\n", root, fb.nfiles, float64(fb.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(rootDir string, dir string, n *sync.WaitGroup, fileSizes chan<- FileSizes) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(rootDir, subdir, n, fileSizes)
		} else {
			fileSizes <- FileSizes{rootDir, entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
