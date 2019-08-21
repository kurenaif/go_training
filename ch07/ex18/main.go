package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	res := "<" + e.Type.Local
	sep := " "
	for _, attr := range e.Attr {
		res += sep + attr.Name.Local + "=\"" + attr.Value + "\""
	}
	res += ">"
	return res
}

func printXML(e *Element, frontSpace string) {
	for _, c := range e.Children {
		switch c := c.(type) {
		case CharData:
			str := strings.TrimSpace(string(c))
			if str != "" {
				fmt.Printf("%s%s\n", frontSpace, str)
			}
		case *Element:
			fmt.Printf("%s%s\n", frontSpace, c.String())
			printXML(c, frontSpace+"  ")
			fmt.Printf("%s</%s>\n", frontSpace, c.Type.Local)
		}
	}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	stack := []*Element{&Element{xml.Name{"root", ""}, nil, []Node{}}} // element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			nextElement := Element{tok.Name, tok.Attr, []Node{}}
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &nextElement)
			stack = append(stack, &nextElement)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			str := fmt.Sprintf("%s", tok)
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(str))
		}
	}

	printXML(stack[0], "")
}
