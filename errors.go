// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import (
	"errors"
	"fmt"
)

// ParseError represents an error occurring during parsing.
type ParseError struct {
	inner error
	msg, Source string
	Line, Col int
}

// Error returns an error message describing the ParseError.
func (err ParseError) Error() string {
	// Set error message format
	format := "bfk.ParseError at line %d, col %d: %s\n"

	// Set default error message
	message := "Invalid program"

	// Set true error message if found
	if  err.msg != "" {
		message = err.msg
	} else if err.inner != nil {
		message = err.inner.Error()
	}

	// Return final message
	return fmt.Sprintf(format, err.Line, err.Col, message)
}

// Unwrap returns the inner error that caused the ParseError.
func (err ParseError) Unwrap() error {
	return err.inner
}

var (
	// ErrBadReader is returned when the io.ReadWriter given
	// to Execute cannot be read from.
	ErrBadReader = errors.New("bfk: Invalid Reader")

	// ErrBadWriter is returned when the io.ReadWriter given
	// to Execute cannot be written to.
	ErrBadWriter = errors.New("bfk: Invalid Writer")

	// ErrBadInput is returned when the io.ReadWriter given
	// to Execute reads an input that the Program cannot
	// understad.
	ErrBadInput = errors.New("bfk: Invalid Input")

	// ErrBadOutput is returned when the io.ReadWriter given
	// to Execute cannot write the output of the Program.
	ErrBadOutput = errors.New("bfk: Invalid Output")
)
