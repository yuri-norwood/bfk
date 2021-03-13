package bfk

var (
	tape_size = 30000
)

type Tape struct {
	pointer int64
	cells   []Cell
}

func (tape *Tape) Right() {
	// Increment the value
	tape.pointer++

	// Check maximum not exceeded
	if tape.pointer >= tape_size {
		tape.pointer = 0
	}
}

func (tape *Tape) Left() {
	// Decrement the value
	tape.pointer--

	// Check minimum not exceeded
	if tape.pointer < 0 {
		tape.pointer = tape_size - 1
	}
}

func (tape *Tape)  Increment() {
	tape.cells[tape.pointer].Increment()
}

func (tape *Tape)  Decrement() {
	tape.cells[tape.pointer].Decrement()
}
