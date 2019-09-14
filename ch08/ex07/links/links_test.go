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
		{"https://exmaple.com", "/index.html"},
		{"https://exmaple.com//////", "/index.html"},
		{"https://exmaple.com/hoge", "/hoge.html"},
		{"https://exmaple.com/hoge/", "/hoge/index.html"},
		{"https://exmaple.com/index.html", "/index.html"},
		{"https://exmaple.com/lib/doc/style.css", "/lib/doc/style.css"},
		{"https://exmaple.com/lib/doc/image.png", "/lib/doc/image.png"},
		{"https://exmaple.com/lib/doc/image.svg", "/lib/doc/image.svg"},
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
