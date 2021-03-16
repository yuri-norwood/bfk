package bfk

import (
	"io"
)

type Program interface {
	IncrementDecrementer
	LeftRighter
	OutputInputter
	io.ReadWriter
}

type program struct {
	tape
	code string
}

func (prgm *program) Read(p []byte) (n int, err error) {
	return 0, errors.New("Unimplemented")
}

func (prgm *program) Write(p []byte) (n int, err error) {
	return 0, errors.New("Unimplemented")
}
