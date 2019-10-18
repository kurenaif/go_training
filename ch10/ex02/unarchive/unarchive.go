package unarchive

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
)

var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

type UnArchveObject interface {
	OpenWithClone(filename string) (UnArchveObject, error)
	Next() (bool, os.FileInfo, *bytes.Buffer, error)
	Close()
}

type format struct {
	name   string
	magic  []byte
	offset int
	uo     UnArchveObject
}

func RegisterFormat(name string, magic []byte, offset int, uo UnArchveObject) {
	formatsMu.Lock()
	defer formatsMu.Unlock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, magic, offset, uo}))
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

// zipがio.raederに対応してないのでfilenameにした
func Open(filename string) (UnArchveObject, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	f := sniff(asReader(file))

	if f.uo == nil {
		return nil, errors.New("unarchive: unknown format")
	}

	return f.uo.OpenWithClone(filename)
}
