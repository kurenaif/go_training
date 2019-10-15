package client

// TODO: インターフェイスでいい感じにエラーハンドリングする

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	conn          net.Conn
	currentDir    string
	addr          string
	port          int
	connectWait   sync.WaitGroup
	connectError  error
	transferWait  sync.WaitGroup
	transferError error
}

var mu sync.Mutex

func (c *Client) SetDir(path string) error {
	path, err := filepath.Abs(path)
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("path not exist")
	}
	if err != nil {
		return err
	}
	c.currentDir = path
	return nil
}

func (c *Client) Dir() (string, error) {
	if c.currentDir == "" {
		err := c.SetDir("./")
		if err != nil {
			return "", err
		}
	}
	return c.currentDir, nil
}

func (c *Client) SetPort(content string) error {
	contents := strings.Split(content, ",")
	if len(contents) != 6 {
		return fmt.Errorf("parse error: PORT requires 6 elements")
	}
	c.addr = fmt.Sprintf("%s.%s.%s.%s", contents[0], contents[1], contents[2], contents[3])

	upper, err := strconv.Atoi(contents[4])
	if err != nil {
		log.Print(err)
		return err
	}
	lower, err := strconv.Atoi(contents[5])
	if err != nil {
		log.Print(err)
		return err
	}
	c.port = upper*256 + lower
	return nil
}

func (c *Client) Addr() string {
	return fmt.Sprintf("%s:%d", c.addr, c.port)
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.Addr())
	if err != nil {
		log.Print(err)
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Close() error {
	err := c.conn.Close()
	return err
}

func (c *Client) ls() (files []os.FileInfo, _ error) {
	dir, err := c.Dir()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (c *Client) List() error {
	err := c.Connect()
	defer c.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	files, err := c.ls()
	if err != nil {
		return err
	}

	res := ""
	for _, file := range files {
		res += file.Name()
		if file.IsDir() {
			res += "/"
		}
		res += "\r\n"
	}
	io.WriteString(c.conn, res)

	return nil
}

func (c *Client) ChangeDirectory(path string) error {
	currentDir, err := c.Dir()
	if err != nil {
		log.Print(err)
		return err
	}
	err = c.SetDir(filepath.Join(currentDir, path))
	if err != nil {
		log.Print(err)
		return fmt.Errorf("451 can't read directory info\r\n")
	}
	return nil
}

func (c *Client) RETR(path string, handleResponse func() error) error {
	err := c.Connect()
	defer c.Close()
	if err != nil {
		return err
	}

	if err := handleResponse(); err != nil {
		return err
	}

	path, err = filepath.Abs(filepath.Join(c.currentDir, path))
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("path not exist")
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) STOR(path string) {
	mu.Lock() //複数writeは心臓に悪い
	defer mu.Unlock()
	// prepare 同時接続数は1と仮定
	c.connectWait = sync.WaitGroup{}
	c.connectWait.Add(1)
	c.transferWait = sync.WaitGroup{}
	c.transferWait.Add(1)
	c.connectError = nil
	c.transferError = nil

	// 接続処理
	err := c.Connect()
	defer c.Close()
	if err != nil {
		c.connectError = err
		return
	}
	c.connectWait.Done()

	// 値を受け取る
	buf, err := ioutil.ReadAll(c.conn)
	if err != nil {
		c.transferError = err
		return
	}

	// safe file
	savePath := ""
	if err != nil {
		c.transferError = err
		return
	}
	if filepath.IsAbs(path) {
		savePath = path
	} else {
		currentDir, err := c.Dir()
		if err != nil {
			c.transferError = err
			return
		}
		savePath = filepath.Join(currentDir, path)
	}

	ioutil.WriteFile(savePath, buf, 0777)

	c.transferWait.Done()
}

func (c *Client) ConnectWait() error {
	c.connectWait.Wait()
	return c.connectError
}

func (c *Client) TransferWait() error {
	c.transferWait.Wait()
	return c.transferError
}

func (c *Client) Pwd() error {
	path, err := c.Dir()
	if err != nil {
		log.Print(err)
		return err
	}
	io.WriteString(c.conn, path)

	return nil
}
