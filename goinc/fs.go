package goinc

import (
	"io"
	"os"
)

var FS FileSystem = DefaultFS{}

type FileSystem interface {
	Open(name string) (File, error)
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt
	Stat() (os.FileInfo, error)
	Fd() uintptr
}

type DefaultFS struct{}

func (DefaultFS) Open(name string) (File, error) {
	return os.Open(name)
}

type StubFile struct{}

func (*StubFile) Close() error {
	return nil
}

func (*StubFile) Read(p []byte) (n int, err error) {
    return 0,nil
}

func (*StubFile) ReadAt(p []byte, off int64) (n int, err error) {
    return 0,nil
}

func (*StubFile) Seek(offset int64, whence int) (int64, error) {
    return 0, nil
}

func (*StubFile) Write(p []byte) (n int, err error) {
    return 0, nil
}

func (*StubFile) WriteAt(p []byte, off int64) (n int, err error) {
    return 0, nil
}

func (*StubFile) Stat() (os.FileInfo, error) {
    return nil, nil
}

func (*StubFile) Fd() uintptr {
    return 0
}