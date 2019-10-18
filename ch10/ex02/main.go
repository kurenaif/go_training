package main

import (
	"fmt"
	"go_training/ch10/ex02/unarchive"
	_ "go_training/ch10/ex02/unarchive/tar"
	_ "go_training/ch10/ex02/unarchive/zip"
	"log"
)

func main() {
	unarchive.ListFormat()
	unTarTest()
	unZipTest()
}

func unArchiveTest(filename string) {
	uo, err := unarchive.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer uo.Close()
	for {
		nxt, fileInfo, bf, err := uo.Next()
		if !nxt {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("filename: ", fileInfo.Name())
		fmt.Println("---content---")
		fmt.Println(string(bf.Bytes()))
	}
}

func unTarTest() {
	unArchiveTest("a.tar")
}

func unZipTest() {
	unArchiveTest("a.zip")
}
