// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import "fmt"

const (
	// errorFormat is the default format for all bfk errors.
	errorFormat = "bfk: %s\n"

	// ErrBadReader is returned when the io.ReadWriter given
	// to Execute cannot be read from.
	ErrBadReader Error = "Invalid Reader"

	// ErrBadWriter is returned when the io.ReadWriter given
	// to Execute cannot be written to.
	ErrBadWriter Error = "Invalid Writer"

	// ErrBadInput is returned when the io.ReadWriter given
	// to Execute reads an input that the Program cannot
	// understad.
	ErrBadInput Error = "Invalid Input"

	// ErrBadOutput is returned when the io.ReadWriter given
	// to Execute cannot write the output of the Program.
	ErrBadOutput Error = "Invalid Output"
)

// ParseError represents an error occurring during parsing.
type ParseError struct {
	inner       error
	msg, Source string
	Line, Col   int
}

// Error represents an error within the bfk package.
type Error string

// Error returns an error message describing the ParseError.
func (err ParseError) Error() string {
	// Set error message format
	format := fmt.Sprintf(errorFormat, "ParseError at line %d, col %d: %s")

	// Set default error message
	message := "Invalid program"

	// Set true error message if found
	if err.msg != "" {
		message = err.msg
	} else if err.inner != nil {
		message = err.inner.Error()
	}

	// Return final message
	return fmt.Sprintf(format, err.Line, err.Col, message)
}

// Error returns an error message describing the Error.
func (err Error) Error() string {
	message := string(err)

	return fmt.Sprintf(errorFormat, message)
}

// Unwrap returns the inner error that caused the ParseError.
func (err ParseError) Unwrap() error {
	return err.inner
}
