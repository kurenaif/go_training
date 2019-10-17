package unarchive

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
)

type format struct {
	name   string
	magic  []byte
	offset int
	decode func(io.Reader) (os.FileInfo, error)
}

var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

func RegisterFormat(name string, magic []byte, offset int, decode func(io.Reader) (os.FileInfo, error)) {
	formatsMu.Lock()
	defer formatsMu.Unlock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, magic, offset, decode}))
}

// 現在登録されているformatを列挙する(デバッグ用)
func ListFormat() {
	formatsMu.Lock()
	defer formatsMu.Unlock()
	formats, _ := atomicFormats.Load().([]format)
	for _, format := range formats {
		fmt.Println(format.name)
	}
}

func sniff(r reader) format {
	formats, _ := atomicFormats.Load().([]format)
	for _, f := range formats {
		b, err := r.Peek(f.offset + len(f.magic))
		if err == nil && f.offset+len(f.magic) <= len(b) && reflect.DeepEqual(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

func asReader(r io.Reader) reader {
	// readerのinterfaceの要件(Peekを持っているかどうか)を満たしているかチェック
	if rr, ok := r.(reader); ok {
		return rr
	}
	// 満たしていなければ、bufioで包んで満たさせる
	return bufio.NewReader(r)
}

func Decode(r io.Reader) (os.FileInfo, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.decode == nil {
		return nil, errors.New("unarchive: unknown format")
	}
	m, err := f.decode(rr)
	return m, err
}
