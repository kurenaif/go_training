package patch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_training/ch04/ex11/ghutil"
	"go_training/ch04/ex11/read"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const PatchIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues/"

type Issue struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type User struct {
	Login string
}

// omitemptyのissueがほしいため、create.goのそれとは別
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

func PatchIssue(token string, number int) error {
	url := PatchIssueURL + strconv.Itoa(number)
	read.ReadIssue(token, number)
	issue, err := CreateIssueBody()
	if err != nil {
		return err
	}
	bodyJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(bodyJSON)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}
