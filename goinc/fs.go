package goinc

import (
    "os"
    "io"
)

var FS FileSystem = FS_t{}

type FileSystem interface {
    Open(name string) (File, error)
}

type File interface {
    io.Closer
    io.Reader
    io.ReaderAt
    io.Seeker
    Stat() (os.FileInfo, error)
    Fd() uintptr
}

type FS_t struct {}

func (FS_t) Open(name string) (File, error) {
    return os.Open(name)
}