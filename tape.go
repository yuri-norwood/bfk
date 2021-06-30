// This is free and unencumbered software released into the public domain.

package bfk

// tape is a private implementor of the memory interface,
// providing access to a collection of cells.
type tape struct {
	cell
	pointer int64
	memory  []cell
	config  *Config
}

// Configure sets the runtime behaviour of a progam running
// on this tape.
func (t *tape) Configure(config Config) {
	t.config = &config
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
		t.memory = append(t.memory, cell{config: t.config})
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
		t.memory = append(t.memory, cell{config: t.config})
	}

	// Update current cell
	t.cell = t.memory[t.pointer]
}
