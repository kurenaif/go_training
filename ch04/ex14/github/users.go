package github

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const UsersURL = "https://api.github.com/users"

var userList = template.Must(template.New("userList").Parse(`
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.ID}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func GetUserPage() (string, error) {
	resp, err := http.Get(UsersURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("GET %s failed: %s", UsersURL, resp.Status)
	}
	var users []*User
	if err = json.NewDecoder(resp.Body).Decode(&users); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()

	var res bytes.Buffer
	if err := userList.Execute(&res, users); err != nil {
		log.Fatal(err)
	}
	return res.String(), nil
}
