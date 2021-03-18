// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

var tapeSize int64 = 30000

type tape struct {
	cell
	pointer int64
	memory  []cell
}

func (t *tape) Right() {
	// Increment the value
	t.pointer++

	// Check maximum not exceeded
	if t.pointer >= tapeSize {
		t.pointer = 0
	}

	// Expand memory
	if t.pointer < int64(len(t.memory)) {
		t.memory = append(t.memory, cell(0))
	}

	// Update current cell
	t.cell = t.memory[t.pointer]
}

func (t *tape) Left() {
	// Decrement the value
	t.pointer--

	// Check minimum not exceeded
	if t.pointer < 0 {
		t.pointer = tapeSize - 1
	}

	// Expand memory
	if t.pointer < int64(len(t.memory)) {
		t.memory = append(t.memory, cell(0))
	}

	// Update current cell
	t.cell = t.memory[t.pointer]
}
