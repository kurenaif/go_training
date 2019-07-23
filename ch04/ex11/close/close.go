package close

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_training/ch04/ex11/read"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Issue struct {
	State string `json:"state"`
}

const PatchIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues/"

func CloseIssue(token string, number int) error {
	url := PatchIssueURL + strconv.Itoa(number)
	read.ReadIssue(token, number)
	issue := Issue{State: "closed"}
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
