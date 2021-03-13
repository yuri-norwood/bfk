package bfk

type Tape struct {
	pointer Pointer
	cells   []Cell
}

func NewTape(size int64) *Tape {
	return &Tape{ cells : make([]Cell, pointer_max) }
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
