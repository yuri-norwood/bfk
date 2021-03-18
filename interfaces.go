// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

// Incrementer represents the ability to increase a cell's value.
type Incrementer interface {
	Increment()
}

// Decrementer represents the ability to decrease a cell's value.
type Decrementer interface {
	Decrement()
}

// IncrementDecrementer represents the ability to both
// increase and decrease a cell's value.
type IncrementDecrementer interface {
	Decrementer
	Incrementer
}

// Lefter represents the ability to move left on a tape.
type Lefter interface {
	Left()
}

// Righer represents the ability to move right on a tape.
type Righter interface {
	Right()
}

// LeftRighter represents the ability to move in both
// directions on a tape.
type LeftRighter interface {
	Lefter
	Righter
}

// Outputter represents the ability to retrieve a cell's value.
type Outputter interface {
	Output() int64
}

// Inputter represents the ability to store a value in a cell.
type Inputter interface {
	Input(int64)
}

// OutputInputter represents the ability to both store and
// retrieve values to and from a cell.
type OutputInputter interface {
	Outputter
	Inputter
}
