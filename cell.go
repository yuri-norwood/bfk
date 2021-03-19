// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

var (
	// cellMax defines the maximum value a cell can contain.
	cellMax int64 = 255

	// cellMin defines the miniumum value a cell can contain.
	cellMin int64 = 0
)

// cell represents a sign value within a memory tape.
type cell int64

// increment increases a cell's value.
func (c *cell) increment() {
	// Increment the value
	value := int64(*c) + 1

	// Check maximum not exceeded
	if value > cellMax {
		value = cellMin
	}

	// Assign new value
	*c = cell(value)
}

// decrement decreases a cell's value.
func (c *cell) decrement() {
	// Decrement the value
	value := int64(*c) - 1

	// Check minimum not exceeded
	if value < cellMin {
		value = cellMax
	}

	// Assign new value
	*c = cell(value)
}

// output retrieves a cell's value.
func (c *cell) output() int64 {
	return int64(*c)
}

// input stores a value in a cell.
func (c *cell) input(value int64) {
	*c = cell(value)
}
