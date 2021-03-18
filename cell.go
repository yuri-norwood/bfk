// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

var (
	cell_max int64 = 255
	cell_min int64 = 0
)

type cell int64

func (c *cell) Increment() {
	// Increment the value
	value := int64(*c) + 1

	// Check maximum not exceeded
	if value > cell_max {
		value = cell_min
	}

	// Assign new value
	*c = cell(value)
}

func (c *cell) Decrement() {
	// Decrement the value
	value := int64(*c) - 1

	// Check minimum not exceeded
	if value < cell_min {
		value = cell_max
	}

	// Assign new value
	*c = cell(value)
}
