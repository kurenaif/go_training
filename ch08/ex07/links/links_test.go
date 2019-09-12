package links

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL2Filepath(t *testing.T) {
	tests := []struct {
		uri  string
		want string
	}{
		{"https://exmaple.com", "index.html"},
		{"https://exmaple.com//////", "index.html"},
		{"https://exmaple.com/hoge", "hoge.html"},
		{"https://exmaple.com/hoge/", "hoge/index.html"},
		{"https://exmaple.com/index.html", "index.html"},
		{"https://exmaple.com/hoge/index.html", "hoge/index.html"},
		{"https://exmaple.com/hoge/fuga.html", "hoge/fuga.html"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("URL2Filepath(%s)", test.uri)
		u, err := url.Parse(test.uri)
		if err != nil {
			t.Errorf("url parse error: %s", err)
		}
		got := URL2Filepath(u)
		if got != test.want {
			t.Errorf("%s = %q want %q", descr, got, test.want)
		}
	}
}
