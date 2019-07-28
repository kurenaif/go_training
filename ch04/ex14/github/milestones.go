package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const MilestoneURL = "https://api.github.com/repos/rust-lang/rust/milestones"

type Milestone struct {
	HTMLURL   string `json:"html_url"`
	Title     string
	Creator   *User
	CreatedAt time.Time `json:"created_at"`
	Number    int
}

//!+template

var milestoneList = template.Must(template.New("milestoneList").Funcs(template.FuncMap{"myTime": myTime}).Parse(`
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Creator</th>
  <th>Title</th>
  <th>Time</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | myTime}}</td>
</tr>
{{end}}
</table>
`))

func GetMilestonePage() (string, error) {
	resp, err := http.Get(MilestoneURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("GET %s failed: %s", MilestoneURL, resp.Status)
	}
	var milestones []*Milestone
	if err = json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()

	var res bytes.Buffer
	if err := milestoneList.Execute(&res, milestones); err != nil {
		log.Fatal(err)
	}
	return res.String(), nil
}

func myTime(t time.Time) string {
	return fmt.Sprint(t)
}
