// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

// Config defines the behaviour of a bfk Program.
type Config struct {
	CellWrap bool
	CellMax  int64
	CellMin  int64
	TapeSize int64
}
