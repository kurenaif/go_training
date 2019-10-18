package tar

import (
	"archive/tar"
	"bytes"
	"fmt"
	"go_training/ch10/ex02/unarchive"
	"io"
	"os"
)

func decode(filename string) (os.FileInfo, *bytes.Buffer, error) {
	fmt.Println("tar decoded!")
	return nil, nil, nil
}

type UntarObject struct {
	filename string
	file     *os.File
	reader   *tar.Reader
}

func (u *UntarObject) OpenWithClone(filename string) (unarchive.UnArchveObject, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	uo := UntarObject{filename, nil, nil}
	tarReader := tar.NewReader(file)
	uo.file = file
	uo.reader = tarReader
	return &uo, nil
}

func (uo *UntarObject) Next() (bool, os.FileInfo, *bytes.Buffer, error) {
	tarHeader, err := uo.reader.Next()
	if err == io.EOF {
		return false, nil, nil, err
	}
	if err != nil {
		return true, nil, nil, err
	}
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, uo.reader); err != nil {
		// 最初のboolは続行フラグで、ここは個々のファイルのエラーを表しているのでtrueを返す。
		return true, nil, nil, err
	}
	return true, tarHeader.FileInfo(), buf, nil
}

func (uo *UntarObject) Close() {
	uo.file.Close()
}

func init() {
	unarchive.RegisterFormat("tar", []byte{0x75, 0x73, 0x74, 0x61, 0x72}, 0x101, &UntarObject{"", nil, nil})
}
