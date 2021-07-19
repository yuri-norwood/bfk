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

const (
	increment = '+'
	decrement = '-'
	input     = ','
	output    = '.'
	left      = '<'
	right     = '>'
	startLoop = '['
	closeLoop = ']'
)

// String returns a string representation of the internal
// state of a program.
func (p *program) String() string {
	return p.code
}

// Execute runs a compiled program instance using the given
// ReadWriter.
func (p *program) Execute(readWriter io.ReadWriter) error {
	line, col := 1, 0

	for _, char := range p.code {
		col++

		switch char {
		case increment:
			p.increment()
		case decrement:
			p.decrement()
		case input:
			// p.input()
			fallthrough
		case output:
			// p.output()
			return ParseError{
				msg:    "IO not yet supported.",
				Source: "program.go",
				Line:   line,
				Col:    col,
			}
		case left:
			p.left()
		case right:
			p.right()
		case startLoop:
			fallthrough
		case closeLoop:
			return ParseError{
				msg:    "Looping not yet supported.",
				Source: "program.go",
				Line:   line,
				Col:    col,
			}
		case '\n':
			line++
			col = 1
		}
	}

	return nil
}
