// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

var (
	tapeSize int64 = 30000
)

type tape struct {
	pointer int64
	cells   []cell
}

func (t *tape) Current() *cell {
	if t.pointer < int64(len(t.cells)) {
		t.cells = append(t.cells, cell(0))
	}

	return &t.cells[t.pointer]
}

func (t *tape) Right() {
	// Increment the value
	t.pointer++

	// Check maximum not exceeded
	if t.pointer >= tapeSize {
		t.pointer = 0
	}
}

func (t *tape) Left() {
	// Decrement the value
	t.pointer--

	// Check minimum not exceeded
	if t.pointer < 0 {
		t.pointer = tapeSize - 1
	}
}

func (t *tape) Increment() {
	t.Current().Increment()
}

func (t *tape) Decrement() {
	t.Current().Decrement()
}

func (t *tape) Output() int64 {
	return t.Current().Output()
}

func (t *tape) Input(value int64) {
	t.Current().Input(value)
}
