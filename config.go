// This is free and unencumbered software released into the public domain.

package bfk

// Config defines the behaviour of a bfk Program.
type Config struct {
	CellWrap   bool
	CellSigned bool
	CellSize   int64
	TapeWrap   bool
	TapeSize   int64
}
