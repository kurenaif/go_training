package split

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	tests := []struct {
		s       string
		sep     string
		wantLen int
		wantArr []string
	}{
		{"a:b:c", ":", 3, []string{"a", "b", "c"}},
		{":::::", ":", 6, []string{"", "", "", "", "", ""}},
		{"", ":", 1, []string{""}},
		{" : ", ":", 2, []string{" ", " "}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("Split(%q, %q)", test.s, test.sep)
		words := strings.Split(test.s, test.sep)
		got := len(words)
		if got != test.wantLen {
			t.Errorf("%s return %d words, want %d", descr, got, test.wantLen)
		}
		if !reflect.DeepEqual(words, test.wantArr) {
			t.Errorf("%s return %q , want %q", descr, words, test.wantArr)
		}
	}
}
