package bfk

import "io"

type Program interface {
	Execute(io.ReadWriter) error
}

type program struct {
	tape
	code string
}

func (p *program) Execute(readWriter io.ReadWriter) error {
	return nil
}

func Parse(text string) Program {
	return &program{code: text}
}
