package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_training/ch04/ex11/ghutil"
	"io/ioutil"
	"net/http"
	"strings"
)

const CreateIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues"

type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

func CreateIssue(token string) error {
	issue, err := CreateIssueBody()
	if err != nil {
		return err
	}
	// make body
	bodyJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", CreateIssueURL, bytes.NewBuffer([]byte(bodyJSON)))
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

func CreateIssueBody() (Issue, error) {
	body := Issue{}

	title, err := ghutil.MessageAndInput("title: ")
	if err != nil {
		return Issue{}, err
	}
	body.Title = title

	body.Body, err = ghutil.MessageAndInput("body: ")
	if err != nil {
		return Issue{}, err
	}

	body.Title = title
	input, err := ghutil.MessageAndInput("asignees(, separated): ")
	if err != nil {
		return Issue{}, err
	}
	assignees := strings.Split(input, ",")
	for _, assignee := range assignees {
		if assignee != "" {
			body.Assignees = append(body.Assignees, assignee)
		}
	}

	input, err = ghutil.MessageAndInput("labels(, separated): ")
	labels := strings.Split(input, ",")
	for _, label := range labels {
		if label != "" {
			body.Labels = append(body.Labels, label)
		}
	}
	return body, nil
}
