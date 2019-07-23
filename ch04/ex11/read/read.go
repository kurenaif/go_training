package read

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const ReadIssueURL = "https://api.github.com/repos/kurenaif/go_training/issues/"

type Issue struct {
	Number int
	Title  string
	Body   string
	User   *User
}

type User struct {
	Login string
}

func ReadIssue(token string, number int) error {
	url := ReadIssueURL + strconv.Itoa(number)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
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

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	fmt.Println("================================================================")
	fmt.Printf("#%04d\t%s\t%s\n", result.Number, result.Title, result.User.Login)
	fmt.Println("================================================================")
	fmt.Printf("%s\n", result.Body)
	fmt.Println("================================================================")
	return nil
}
