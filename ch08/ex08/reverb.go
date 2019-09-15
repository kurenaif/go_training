package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	defer c.Close()
	ch := make(chan string)
	eof := make(chan struct{})

	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
		eof <- struct{}{}
	}()

	for {
		select {
		case <-time.After(10 * time.Second):
			log.Print("connection closed because of time out")
			fmt.Fprintln(c, "time out")
			return
		case <-eof:
			log.Print("input end")
			break
		case s := <-ch:
			fmt.Println(s)
			go echo(c, s, 1*time.Second)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
