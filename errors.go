// This is free and unencumbered software released into the public domain.

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
	// understand.
	ErrBadInput Error = "Invalid Input"

	// ErrBadOutput is returned when the io.ReadWriter given
	// to Execute cannot write the output of the Program.
	ErrBadOutput Error = "Invalid Output"
)

// parseError represents an error occurring during parsing.
type parseError struct {
	inner       error
	msg, source string
	line, col   int
}

// Error represents an error within the bfk package.
type Error string

// Error returns an error message describing the parseError.
func (err parseError) Error() string {
	// Set error message format
	format := fmt.Sprintf(errorFormat, "ParseError in %s at line %d, col %d: %s")

	// Set default source file
	source := "package bfk"

	// Set default line number
	line := -1

	// Set default column number
	col := -1

	// Set default error message
	message := "Invalid program"

	// Get error source if given
	if err.source != "" {
		source = err.source
	}

	// Get error line number if given
	if err.line > 0 {
		line = err.line
	}

	// Get error coloumn number if given
	if err.col > 0 {
		col = err.col
	}

	// Get error message if found
	if err.msg != "" {
		message = err.msg
	} else if err.inner != nil {
		message = err.inner.Error()
	}

	// Return final message
	return fmt.Sprintf(format, source, line, col, message)
}

// Error returns an error message describing the Error.
func (err Error) Error() string {
	message := string(err)

	return fmt.Sprintf(errorFormat, message)
}

// Unwrap returns the inner error that caused the parseError.
func (err parseError) Unwrap() error {
	return err.inner
}
