// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import (
	"errors"
	"io"
)

// Program provides external access to compiled Brainfuck
// program to execute.
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

// Parse provides a safe way of compiling a Brainfuck program
// and creating an external Program to access and execute.
func Parse(text string) Program {
	return &program{code: text}
}
