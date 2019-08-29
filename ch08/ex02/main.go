package main

import (
	"bufio"
	"fmt"
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

func string2Request(str string) (Request, error) {
	code, err := strconv.Atoi(str[0:3])
	if err != nil {
		return Request{}, fmt.Errorf("request parse error: %s is not number: %s", str[0:3], err)
	}
	return Request{code, str[4:len(str)]}, nil
}

func (r Request) String() string {
	return fmt.Sprintf("%d %s\r\n", r.code, r.message)
}

func writeRequest(c net.Conn, req Request) {
	io.WriteString(c, req.String())
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	_, err := io.WriteString(c, "220 (kurenaiftp)\r\n")
	if err != nil {
		return
	}
	requestScanner := bufio.NewScanner(c)
	for requestScanner.Scan() {
		request := requestScanner.Text()
		fmt.Println(request)
		strs := strings.Split(request, " ")
		if strs[0] == "USER" {
			writeRequest(c, Request{331, "Please specify the password"})
		}
		if strs[0] == "PASS" {
			writeRequest(c, Request{230, "Login successful"})
		}
		if strs[0] == "SYST" {
			writeRequest(c, Request{215, "UNIX Type: L8"})
		}
	}
}
