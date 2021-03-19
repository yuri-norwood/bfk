// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import "io"

// program is a private implementor of the Program interface,
// to be returned by the Parse constructor.
type program struct {
	tape
	code string
}

// Execute runs a compiled program instance using the given
// ReadWriter.
func (p *program) Execute(readWriter io.ReadWriter) error {
	return nil
}

// Parse provides a safe way of compiling a Brainfuck program
// and creating an external Program to access and execute.
func Parse(text string) Program {
	return &program{code: text}
}
