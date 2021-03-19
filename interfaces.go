// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

import "io"

// incrementer represents the ability to increase a cell's value.
type incrementer interface {
	increment()
}

// decrementer represents the ability to decrease a cell's value.
type decrementer interface {
	decrement()
}

// incrementDecrementer represents the ability to both
// increase and decrease a cell's value.
type incrementDecrementer interface {
	decrementer
	incrementer
}

// lefter represents the ability to move left on a tape.
type lefter interface {
	left()
}

// righter represents the ability to move right on a tape.
type righter interface {
	right()
}

// leftRighter represents the ability to move in both
// directions on a tape.
type leftRighter interface {
	lefter
	righter
}

// outputter represents the ability to retrieve a cell's value.
type outputter interface {
	output() int64
}

// inputter represents the ability to store a value in a cell.
type inputter interface {
	input(int64)
}

// outputInputter represents the ability to both store and
// retrieve values to and from a cell.
type outputInputter interface {
	outputter
	inputter
}

// memory represents the behaviour of a tape of cells, allowing
// input and output, moving across the tape, and modifying the
// values of the cells.
type memory interface {
	incrementDecrementer
	leftRighter
	outputInputter
}

// Program provides external access to compiled Brainfuck
// program to execute.
type Program interface {
	memory
	Execute(io.ReadWriter) error
}
