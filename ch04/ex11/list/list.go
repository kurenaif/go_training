package list

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ListIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues"

type Issue struct {
	Number int
	Title  string
	User   *User
}

type User struct {
	Login string
}

func ListIssue(token string) error {
	req, err := http.NewRequest("GET", ListIssueURL, nil)
	if err != nil {
		return err
	}
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("read issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	// b, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	for _, issue := range result {
		fmt.Printf("#%04d, %50s, %s\n", issue.Number, issue.Title, issue.User.Login)
	}
	return nil
}
