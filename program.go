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
type operation int

const (
	increment operation = iota
	decrement operation
	input     operation
	output    operation
	left      operation
	right     operation
	startLoop operation
	closeLoop operation
)

// String returns a string representation of the internal
// state of a program.
func (p *program) String() string {
	return p.code
}

// Execute runs a compiled program instance using the given
// ReadWriter.
func (p *program) Execute(readWriter io.ReadWriter) error {
	operations := make([]operation, 0, len(p.code))

	for _, char := range p.code {
		switch (char)
		{
		case "+":
			operations = append(operations, increment)
		case "-":
			operations = append(operations, decrement)
		case ",":
			operations = append(operations, input)
		case ".":
			operations = append(operations, output)
		case "<":
			operations = append(operations, left)
		case ">":
			operations = append(operations, right)
		case "[":
			operations = append(operations, startLoop)
		case "]":
			operations = append(operations, closeLoop)
		}
	}

	return nil
}
