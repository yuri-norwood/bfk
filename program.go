// This is free and unencumbered software released into the public domain.

package bfk

import "io"

// program is a private implementor of the Program interface,
// to be returned by the Parse constructor.
type program struct {
	tape
	code   string
	config Config
}

// operation is a private value representing a BF operation.
type operation string

const (
	increment operation = "+"
	decrement operation = "-"
	input     operation = ","
	output    operation = "."
	left      operation = "<"
	right     operation = ">"
	startLoop operation = "["
	closeLoop operation = "]"
)

// String returns a string representation of the internal
// state of a program.
func (p *program) String() string {
	return p.code
}

// Execute runs a compiled program instance using the given
// ReadWriter.
func (p *program) Execute(readWriter io.ReadWriter) error {
	for _, char := range p.code {
		switch (char)
		{
		case increment:
		case decrement:
		case input:
		case output:
		case left:
		case right:
		case startLoop:
		case closeLoop:
		}
	}

	return nil
}
