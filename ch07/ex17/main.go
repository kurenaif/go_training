package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Element struct {
	tagName   string
	attribute map[string]string // Local: Value
}

func (e Element) String() string {
	res := "<" + e.tagName
	sep := " "
	for local, value := range e.attribute {
		res += sep + local + "=\"" + value + "\""
	}
	res += ">"
	return res
}

// タグ名一致 & tar ∈ e
func (e Element) include(tar Element) bool {
	if e.tagName != tar.tagName {
		return false
	}
	if len(e.attribute) < len(tar.attribute) {
		return false
	}
	for k, ev := range e.attribute {
		tarV, ok := tar.attribute[k]
		if !ok {
			continue
		}
		if ev != tarV {
			return false
		}
	}
	return true
}

func ElementsJoin(es []Element, sep string) string {
	s := ""
	res := ""
	for _, e := range es {
		res += s + e.String()
		s = sep
	}
	return res
}

func parseArgs() ([]Element, error) {
	elements := []Element{}
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") {
			LocalValue := strings.Split(arg, "=")
			if len(LocalValue) >= 3 {
				return nil, fmt.Errorf("Onle one \"=\" is allowed (given: %s)", arg)
			}
			if len(elements) < 1 {
				return nil, fmt.Errorf("elementName is required before attributes")
			}
			elements[len(elements)-1].attribute[LocalValue[0]] = LocalValue[1]
		} else {
			elements = append(elements, Element{arg, map[string]string{}})
		}
	}
	return elements, nil
}

func main() {
	argElements, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element // stack of element names
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
			attributes := map[string]string{}
			for _, attr := range tok.Attr {
				attributes[attr.Name.Local] = attr.Value
			}
			stack = append(stack, Element{tok.Name.Local, attributes}) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, argElements) {
				fmt.Printf("%s: %s\n", ElementsJoin(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []Element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].include(y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
