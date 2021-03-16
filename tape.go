package bfk

var (
	tape_size = 30000
)

type tape struct {
	pointer int64
	cells   []Cell
}

func (t *tape) Right() {
	// Increment the value
	t.pointer++

	// Check maximum not exceeded
	if t.pointer >= tape_size {
		t.pointer = 0
	}
}

func (t *tape) Left() {
	// Decrement the value
	t.pointer--

	// Check minimum not exceeded
	if t.pointer < 0 {
		t.pointer = tape_size - 1
	}
}

func (t *Tape) Increment() {
	t.Current().Increment()
}

func (t *Tape) Decrement() {
	t.Current().Decrement()
}

func (t *Tape) Current() *cell {
	return t.cells[t.pointer]
}
