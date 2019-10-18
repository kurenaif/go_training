package zip

import (
	"archive/zip"
	"bytes"
	"go_training/ch10/ex02/unarchive"
	"io"
	"log"
	"os"
)

type UnzipObject struct {
	filename string
	reader   *zip.ReadCloser
	index    int
}

func (u *UnzipObject) OpenWithClone(filename string) (unarchive.UnArchveObject, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	uo := UnzipObject{filename, nil, 0}
	uo.index = 0
	uo.reader = r
	return &uo, nil
}

func (uo *UnzipObject) Next() (bool, os.FileInfo, *bytes.Buffer, error) {
	if uo.index >= len(uo.reader.File) {
		return false, nil, nil, nil
	}
	f := uo.reader.File[uo.index]
	r, err := f.Open()
	if err != nil {
		// 最初のboolは続行フラグで、ここは個々のファイルのエラーを表しているのでtrueを返す。
		return true, nil, nil, err
	}
	defer r.Close()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		log.Fatal(err)
	}
	uo.index += 1
	return true, f.FileInfo(), buf, nil
}

func (uo *UnzipObject) Close() {
	uo.reader.Close()
}

func init() {
	unarchive.RegisterFormat("zip", []byte{0x50, 0x4B}, 0x0, &UnzipObject{"", nil, 0})
}
