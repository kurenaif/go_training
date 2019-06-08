package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args   []string
		inputs []string
		want   string
	}{
		{[]string{"a.txt", "b.txt"}, nil, "(sum: 2 ):\n\ta.txt\t1\n\tb.txt\t1\nhello(sum: 3 ):\nhello\ta.txt\t3\nhollow(sum: 2 ):\nhollow\ta.txt\t1\nhollow\tb.txt\t1\ntext(sum: 4 ):\ntext\tb.txt\t2\ntext\ta.txt\t2\n"},
		// {[]string{}, []string{"test", "test"}, "test(sum: 2 ):\ntest\t|0\t2\n"},
	}

	for _, test := range tests {
		descr := ""
		// if test.inputs == nil {
		// descr = fmt.Sprintf("outputCountLines(%q)", test.args)
		// } else {
		// 	descr = fmt.Sprintf("outputCountLines(%q), std. inputs:(%q)", test.args, test.inputs)
		// }
		out = new(bytes.Buffer)
		outputCountLines(test.args)
		// if len(test.args) == 0 {
		// 	outputCountLines(test.args)
		// } else {
		// 	StubIO(strings.Join(test.inputs, "\n"), func() {
		// 		outputCountLines(test.args)
		// 	})
		// }
		got := out.(*bytes.Buffer).String()
		fmt.Println(got)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

// ref: https://qiita.com/nak1114/items/a2a9b036097ea120b12b
// StubIO stubs Stdin Stdout Stderr in 'fn'.return Stdout and Stderr
func StubIO(inbuf string, fn func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()
	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	inw.Write([]byte(inbuf))
	inw.Close()
	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw
	fn()
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	outbuf, _ := ioutil.ReadAll(outr)
	errw.Close()
	errbuf, _ := ioutil.ReadAll(errr)

	return string(outbuf), string(errbuf)

}
