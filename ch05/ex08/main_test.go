package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	var tests = []struct {
		content string
		want    string
	}{
		{
			`<html>
<!-- mycomment -->
<head>
		</head>
<body>
				<img src="image.jpg" alt="kurenaif"></img>
  <p> hello     world</p>
  <div> <a href="link.html"> link </a> </div>
</body>
</html>`,

			`<html>
  <!-- mycomment -->
  <head>
  </head>
  <body>
    <img src="image.jpg" alt="kurenaif"/>
    <p>
      hello     world
    </p>
    <div>
      <a href="link.html">
        link
      </a>
    </div>
  </body>
</html>
`,
		},
	}

	for _, test := range tests {
		// descr := fmt.Sprintf("forEachNode(%s, startElement, endElement)", test.content)
		doc, err := html.Parse(strings.NewReader(test.content))
		if err != nil {
			t.Error(err)
		}
		out = new(bytes.Buffer)
		forEachNode(doc, startElement, endElement)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("\n-----------------got-----------------\n%s\n-----------------want-----------------\n%s", got, test.want)
		}
	}
}
