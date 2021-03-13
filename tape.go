package bfk

var (
	tape_size = 30000
)

type Tape struct {
	pointer Pointer
	cells   []Cell
}

func NewTape() *Tape {
	return &Tape{ cells : make([]Cell, tape_size) }
}

func (tape *Tape) Left() {
	tape.pointer.Left()
}

func (tape *Tape) Right() {
	tape.pointer.Right()
}

func (tape *Tape)  Increment() {
	tape.cells[tape.pointer].Increment()
}

func (tape *Tape)  Decrement() {
	tape.cells[tape.pointer].Decrement()
}
