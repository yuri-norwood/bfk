// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

// tape is a private implementor of the memory interface,
// providing access to a collection of cells.
type tape struct {
	cell
	pointer int64
	memory  []cell
	config  *Config
}

// right moves the pointer right on a tape.
func (t *tape) right() {
	// Increment the value
	t.pointer++

	// Check maximum not exceeded
	if t.pointer >= t.config.TapeSize {
		t.pointer = 0
	}

	// Expand memory
	if t.pointer < int64(len(t.memory)) {
		c = cell{value: 0}
		c.config = t.config

		t.memory = append(t.memory, c)
	}

	// Update current cell
	t.cell = t.memory[t.pointer]
}

// left moves the pointer left on a tape.
func (t *tape) left() {
	// Decrement the value
	t.pointer--

	// Check minimum not exceeded
	if t.pointer < 0 {
		t.pointer = t.config.TapeSize - 1
	}

	// Expand memory
	if t.pointer < int64(len(t.memory)) {
		c = cell{value: 0}
		c.config = t.config

		t.memory = append(t.memory, c)
	}

	// Update current cell
	t.cell = t.memory[t.pointer]
}
