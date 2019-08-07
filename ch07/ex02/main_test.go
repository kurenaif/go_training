package main

import (
	"os"
	"testing"
)

func TestCounter(t *testing.T) {
	writer, cnt := CountingWriter(os.Stdout) //初期は0
	if *cnt != 0 {
		t.Errorf("cnt want 0")
	}

	writer.Write([]byte("")) //0バイト書き込み
	if *cnt != 0 {
		t.Errorf("cnt want 0")
	}

	writer.Write([]byte("Hello\n"))
	if *cnt != 6 {
		t.Errorf("cnt want 6")
	}

	writer.Write([]byte("Hello\n")) //加算
	if *cnt != 12 {
		t.Errorf("cnt want 12")
	}

	writer2, cnt2 := CountingWriter(os.Stderr) // cnt2には引き継がれない
	writer2.Write([]byte("Hello\n"))
	if *cnt2 != 6 {
		t.Errorf("cnt want 6")
	}

	*cnt = 0 // リセット可能
	writer.Write([]byte("Hello\n"))
	if *cnt != 6 {
		t.Errorf("cnt want 6")
	}
}
