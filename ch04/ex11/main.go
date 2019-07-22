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

func openEditor(editor string) (string, error) {
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
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	resByte, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(resByte), "\n"), nil
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

func messageAndInput(message string) (string, error) {
	fmt.Print(message)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	res := s.Text()
	if res == "EDITOR" {
		r, err := openEditor(os.Getenv("EDITOR"))
		if err != nil {
			return "", err
		}
		res = string(r)
	}
	return res, nil
}

func createIssueBody() (CreateIssue, error) {
	body := CreateIssue{}

	title, err := messageAndInput("title: ")
	if err != nil {
		return CreateIssue{}, err
	}
	body.Title = title

	body.Body, err = messageAndInput("body: ")
	if err != nil {
		return CreateIssue{}, err
	}

	body.Title = title
	input, err := messageAndInput("asignees(, separated): ")
	if err != nil {
		return CreateIssue{}, err
	}
	assignees := strings.Split(input, ",")
	for _, assignee := range assignees {
		if assignee != "" {
			body.Assignees = append(body.Assignees, assignee)
		}
	}

	input, err = messageAndInput("labels(, separated): ")
	labels := strings.Split(input, ",")
	for _, label := range labels {
		if label != "" {
			body.Labels = append(body.Labels, label)
		}
	}
	return body, nil
}

func main() {
	// editor := os.Getenv("EDITOR")
	// if editor == "" {
	// 	editor = "vim"
	// }
	// res, _ := openEditor(editor)
	// fmt.Println(string(res))
	s, err := messageAndInput("[create]>")
	if err != nil {
		log.Fatal(err)
	}

	if s[:1] == "c" {
		body, err := createIssueBody()
		if err != nil {
			log.Fatal(err)
		}

		err = createIssue(os.Args[1], body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
