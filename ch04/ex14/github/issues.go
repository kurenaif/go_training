package github

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const IssueURL = "https://api.github.com/repos/rust-lang/rust/issues"

type Issue struct {
	HTMLURL string `json:"html_url"`
	Title   string
	Body    string
	User    *User
	Asignee *User
	Number  int
}

type User struct {
	ID      int
	Login   string
	HTMLURL string `json:"html_url"`
}

var issueList = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
  <th>Title</th>
  <th>Body</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.Body}}</td>
</tr>
{{end}}
</table>
`))

func GetIssuePage() (string, error) {
	resp, err := http.Get(IssueURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("GET %s failed: %s", IssueURL, resp.Status)
	}
	var issues []*Issue
	if err = json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()

	var res bytes.Buffer
	if err := issueList.Execute(&res, issues); err != nil {
		log.Fatal(err)
	}
	return res.String(), nil
}
