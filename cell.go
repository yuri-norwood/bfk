// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

var (
	cellMax int64 = 255
	cellMin int64 = 0
)

type cell int64

func (c *cell) Increment() {
	// Increment the value
	value := int64(*c) + 1

	// Check maximum not exceeded
	if value > cellMax {
		value = cellMin
	}

	// Assign new value
	*c = cell(value)
}

func (c *cell) Decrement() {
	// Decrement the value
	value := int64(*c) - 1

	// Check minimum not exceeded
	if value < cellMin {
		value = cellMax
	}

	// Assign new value
	*c = cell(value)
}

func (c *cell) Output() int64 {
	return int64(*c)
}

func (c *cell) Input(value int64) {
	*c = cell(value)
}
