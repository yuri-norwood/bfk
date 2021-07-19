// This is free and unencumbered software released into the public domain.

package bfk

import "fmt"

// errorFormat is the default format for all bfk errors.
const errorFormat = "bfk: %s\n"

var (
	// ErrBadReader is returned when the io.ReadWriter given
	// to Execute cannot be read from.
	ErrBadReader = fmt.Errorf(errorFormat, "Invalid Reader")

	// ErrBadWriter is returned when the io.ReadWriter given
	// to Execute cannot be written to.
	ErrBadWriter = fmt.Errorf(errorFormat, "Invalid Writer")

	// ErrBadInput is returned when the io.ReadWriter given
	// to Execute reads an input that the Program cannot
	// understand.
	ErrBadInput = fmt.Errorf(errorFormat, "Invalid Input")

	// ErrBadOutput is returned when the io.ReadWriter given
	// to Execute cannot write the output of the Program.
	ErrBadOutput = fmt.Errorf(errorFormat, "Invalid Output")
)

// ParseError represents an error occurring during parsing.
type ParseError struct {
	inner       error
	msg, Source string
	Line, Col   int
}

// Error returns an error message describing the ParseError.
func (err ParseError) Error() string {
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
	if err.Source != "" {
		source = err.Source
	}

	// Get error line number if given
	if err.Line > 0 {
		line = err.Line
	}

	// Get error coloumn number if given
	if err.Col > 0 {
		col = err.Col
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

// Unwrap returns the inner error that caused the ParseError.
func (err ParseError) Unwrap() error {
	return err.inner
}
