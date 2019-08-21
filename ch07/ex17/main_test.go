package main

import (
	"fmt"
	"testing"
)

func TestInclude(t *testing.T) {
	tests := []struct {
		lhs  Element
		rhs  Element
		want bool
	}{
		{Element{}, Element{}, true},
		{Element{}, Element{"tag", map[string]string{"id": "3"}}, false},
		{Element{"tag", map[string]string{"id": "3"}}, Element{"tag", map[string]string{"id": "3"}}, true},
		{Element{"tag", map[string]string{"id": "3"}}, Element{"tag", map[string]string{}}, true},
		{Element{"tag", map[string]string{}}, Element{"tag", map[string]string{"id": "3"}}, false},
		{Element{"tag", map[string]string{"id": "3", "name": "hoge"}}, Element{"tag", map[string]string{"name": "hoge"}}, true},
		{Element{"tag", map[string]string{"name": "hoge"}}, Element{"tag", map[string]string{"id": "3", "name": "hoge"}}, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(%v).include(%v)", test.lhs, test.rhs)
		got := test.lhs.include(test.rhs)
		if got != test.want {
			t.Errorf("%s = %t, want %t", descr, got, test.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		e    Element
		want string
	}{
		{Element{"tag", map[string]string{}}, "<tag>"},
		{Element{"tag", map[string]string{"id": "3"}}, "<tag id=\"3\">"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(%v).String()", test.e)
		got := test.e.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}
