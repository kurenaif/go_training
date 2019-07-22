package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type CreateIssue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

const CreateIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues"

func openEditor(editor string) ([]byte, error) {
	path := os.TempDir() + "/temp"
	tempF, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	tempF.Close()

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return []byte{}, err
	}
	err = cmd.Wait()
	if err != nil {
		return []byte{}, err
	}

	readF, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	res, err := ioutil.ReadAll(readF)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}

func createIssue(token string, issue CreateIssue) error {
	// make body
	body, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", body)

	req, err := http.NewRequest("POST", CreateIssueURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}

func messageAndInput(message string) string {
	fmt.Print(message)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}

func createIssueBody() (body CreateIssue) {
	body.Title = messageAndInput("title: ")
	body.Body = messageAndInput("body: ")
	assignees := strings.Split(messageAndInput("asignees(, separated): "), ",")
	labels := strings.Split(messageAndInput("labels(, separated): "), ",")
	for _, assignee := range assignees {
		if assignee != "" {
			body.Assignees = append(body.Assignees, assignee)
		}
	}
	for _, label := range labels {
		if label != "" {
			body.Labels = append(body.Labels, label)
		}
	}
	return
}

func main() {
	s := messageAndInput("[create]>")
	if s[:1] == "c" {
		err := createIssue(os.Args[1], createIssueBody())
		if err != nil {
			log.Fatal(err)
		}
	}
}
