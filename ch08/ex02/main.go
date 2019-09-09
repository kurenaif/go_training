package main

import (
	"bufio"
	"fmt"
	"go_training/ch08/ex02/client"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	code    int
	message string
}

func String2Request(str string) (Request, error) {
	code, err := strconv.Atoi(str[0:3])
	if err != nil {
		return Request{}, fmt.Errorf("request parse error: %s is not number: %s", str[0:3], err)
	}
	return Request{code, str[4:len(str)]}, nil
}

func (r Request) String() string {
	return fmt.Sprintf("%d %s\r\n", r.code, r.message)
}

func WriteRequest(c net.Conn, req Request) error {
	_, err := io.WriteString(c, req.String())
	if err != nil {
		log.Print(err)
	}
	return err
}

func main() {
	listener, err := net.Listen("tcp", "localhost:2121")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	client := client.Client{}
	_, err := io.WriteString(c, "220 (kurenaiftp)\r\n")
	if err != nil {
		return
	}
	requestScanner := bufio.NewScanner(c)
	for requestScanner.Scan() {
		request := requestScanner.Text()
		fmt.Println(request)
		strs := strings.Split(request, " ")
		switch strs[0] {
		case "USER":
			WriteRequest(c, Request{331, "Please specify the password"})
		case "PASS":
			WriteRequest(c, Request{230, "Login successful"})
		case "SYST":
			WriteRequest(c, Request{215, "UNIX Type: L8"})
		case "QUIT":
			WriteRequest(c, Request{221, "Goodbye."})
		case "PORT":
			if len(strs) < 2 {
				log.Print("PORT request requires space splitted message")
				WriteRequest(c, Request{500, "PORT request requires space splitted message"})
				continue
			}
			err := client.SetPort(strs[1])
			if err != nil {
				WriteRequest(c, Request{425, "data connection can't open"})
				log.Print(err)
				continue
			}
			WriteRequest(c, Request{200, "PORT command successful. Consider using PASV."})
		case "LIST":
			WriteRequest(c, Request{150, "Here comes the directory listing."})
			err := client.List()
			if err != nil {
				WriteRequest(c, Request{550, "Failed get directory"})
				continue
			}
			WriteRequest(c, Request{226, "Directory send ok"})
		case "CWD":
			if len(strs) < 2 {
				log.Print("CWD request requires space splitted message")
				WriteRequest(c, Request{500, "CWD request requires space splitted message"})
				continue
			}
			err := client.ChangeDirectory(strs[1])
			if err != nil {
				WriteRequest(c, Request{550, "Failed to change directory."})
				continue
			}
			WriteRequest(c, Request{250, "Directory successfully changed."})
		case "TYPE":
			if len(strs) < 2 {
				log.Print("TYPE request requires space splitted message")
				WriteRequest(c, Request{500, "TYPE request requires space splitted message"})
				continue
			}
			if strs[1] == "I" {
				// TODO: implementatin
				WriteRequest(c, Request{200, "Switching to Binary mode"})
			} else {
				WriteRequest(c, Request{504, "this parameter is not implmeneted"})
			}
		case "RETR":
			if len(strs) < 2 {
				log.Print("RETR request requires space splitted message")
				WriteRequest(c, Request{500, "RETR request requires space splitted message"})
				continue
			}
			err := client.RETR(strs[1], func() error {
				return WriteRequest(c, Request{150, fmt.Sprintf("Opening BINARY mode data connection for %s", strs[1])})
			})
			if err != nil {
				log.Print(err)
				WriteRequest(c, Request{550, "failed RETR"})
			}
			WriteRequest(c, Request{226, "Transfer complete"})
		case "STOR":
			if len(strs) < 2 {
				log.Print("STOR request requires space splitted message")
				WriteRequest(c, Request{500, "STOR request requires space splitted message"})
				continue
			}
			go client.STOR(strs[1])
			if err := client.ConnectWait(); err != nil {
				log.Print(err)
				WriteRequest(c, Request{550, "failed STOR"})
				continue
			}
			WriteRequest(c, Request{150, "Ok to send data."})
			if err := client.TransferWait(); err != nil {
				log.Print(err)
				WriteRequest(c, Request{550, "failed STOR"})
				continue
			}
			WriteRequest(c, Request{226, "Transfer complete"})
		default:
			WriteRequest(c, Request{504, "this command is not implmeneted"})
		}
	}
}
