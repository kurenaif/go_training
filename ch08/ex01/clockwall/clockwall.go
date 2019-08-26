package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type ClearWriter struct {
	Locate string
	Host   string
	Time   string
}

func (c *ClearWriter) Write(data []byte) (n int, err error) {
	c.Time = string(data)
	return len(data), err
}

func parseArgs() []ClearWriter {
	cw := []ClearWriter{}
	for i := 1; i < len(os.Args); i++ {
		s := os.Args[i]
		ss := strings.Split(s, "=")
		if len(ss) != 2 {
			log.Fatal("invalid format (equal must be one)", s)
		}
		cw = append(cw, ClearWriter{ss[0], ss[1], ""})
	}
	return cw
}

func main() {
	cws := parseArgs()
	fmt.Println(cws)

	for i := range cws {
		cw := &cws[i]
		fmt.Println(cw.Host)
		conn, err := net.Dial("tcp", cw.Host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(cw, conn)
	}

	for {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		for _, cw := range cws {
			if cw.Time != "" {
				fmt.Printf("%s: %s", cw.Locate, cw.Time)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
