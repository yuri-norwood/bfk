package bfk

var (
	tape_size int64 = 30000
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

func (t *tape) Increment() {
	t.Current().Increment()
}

func (t *tape) Decrement() {
	t.Current().Decrement()
}

func (t *tape) Output() int64 {
	return int64(*t.Current())
}

func (t *tape) Input(value int64) {
	*t.Current() = cell(value)
}
