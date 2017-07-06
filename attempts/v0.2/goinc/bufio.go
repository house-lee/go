package goinc

import (
	"bufio"
	"io"
)

var BufIO BufferIO = DefaultBufIO{}

type BufferIO interface {
	NewReader(rd io.Reader) BufIOReader
}

type DefaultBufIO struct{}

type BufIOReader interface {
	Discard(n int) (discarded int, err error)
	Peek(n int) ([]byte, error)
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadLine() (line []byte, isPrefix bool, err error)
	ReadRune() (r rune, size int, err error)
	ReadSlice(delim byte) (line []byte, err error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}

func (DefaultBufIO) NewReader(rd io.Reader) BufIOReader {
	return bufio.NewReader(rd)
}

type StubBufIO struct{}

func (*StubBufIO) Discard(n int) (discarded int, err error)          { return }
func (*StubBufIO) Peek(n int) ([]byte, error)                        { return nil, nil }
func (*StubBufIO) Read(p []byte) (n int, err error)                  { return }
func (*StubBufIO) ReadByte() (byte, error)                           { return 0, nil }
func (*StubBufIO) ReadBytes(delim byte) ([]byte, error)              { return nil, nil }
func (*StubBufIO) ReadLine() (line []byte, isPrefix bool, err error) { return }
func (*StubBufIO) ReadRune() (r rune, size int, err error)           { return }
func (*StubBufIO) ReadSlice(delim byte) (line []byte, err error)     { return }
func (*StubBufIO) ReadString(delim byte) (string, error)             { return "", nil }
func (*StubBufIO) Reset(r io.Reader)                                 {}
func (*StubBufIO) UnreadByte() error                                 { return nil }
func (*StubBufIO) UnreadRune() error                                 { return nil }
func (*StubBufIO) WriteTo(w io.Writer) (n int64, err error)          { return }
