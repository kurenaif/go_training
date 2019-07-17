package main

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"HelloWorld", "dlroWolleH"}, // 文字数の偶奇チェック
		{"Hello World", "dlroW olleH"},
		{"😥ello世界", "界世olle😥"},
		{"😥ell世界", "界世lle😥"},
		{"😥e→😃世:)", "):世😃→e😥"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverse([]byte(%s))", test.s)
		bs := []byte(test.s)
		reverse(bs)
		got := string(bs)
		if got != test.want {
			t.Errorf("%s = %v want %v\n", descr, got, test.want)
		}
	}
}
