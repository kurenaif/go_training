package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var tests = []struct {
		content string
		id      string
		want    string
	}{
		{`<html>
  <!-- mycomment -->
  <head>
  </head>
  <body>
    <img src="image.jpg" alt="kurenaif" id="findThis!"/>
    <p>
      hello     world
    </p>
    <div>
      <a href="link.html" id="findThis!">
        link
      </a>
    </div>
  </body>
</html>
`, // note: ２つ目のfindThisは表示しない
			"findThis!",
			`<img src="image.jpg" alt="kurenaif" id="findThis!">`, // 最後のslashはnodeの情報に含まれていないので表示しない
		},
	}

	for _, test := range tests {
		// descr := fmt.Sprintf("forEachNode(%s, startElement, endElement)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		out = new(bytes.Buffer)
		got := ElementByID(doc, test.id)

		gotString := fmt.Sprintf("<%s", got.Data)
		for _, attr := range got.Attr {
			gotString += fmt.Sprintf(" %s=%q", attr.Key, attr.Val)
		}
		gotString += fmt.Sprintf(">")

		if gotString != test.want {
			t.Errorf("got = %s \n want = %s", gotString, test.want)
		}
	}
}
