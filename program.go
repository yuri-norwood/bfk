// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import "io"

type program struct {
	tape
	code string
}

func (p *program) Execute(readWriter io.ReadWriter) error {
	return nil
}

// Parse provides a safe way of compiling a Brainfuck program
// and creating an external Program to access and execute.
func Parse(text string) Program {
	return &program{code: text}
}
