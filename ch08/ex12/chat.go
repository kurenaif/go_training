package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster

type clientData struct {
	userName string
	client   chan string
}

var (
	entering = make(chan clientData)
	leaving  = make(chan clientData)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[clientData]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.client <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			message := "active members are: "
			for c := range clients {
				message += c.userName + " "
			}
			cli.client <- message

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.client)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- clientData{who, ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- clientData{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
