// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

package bfk

// cell represents a sign value within a memory tape.
type cell struct {
	value int64
	config *Config
}

// increment increases a cell's value.
func (c *cell) increment() {
	// Increment the value
	value := *c.value + 1
	
	// Get limit values
	max = *(*c.config).CellMax
	min = *(*c.config).CellMin

	// Check maximum not exceeded
	if value > max {
		value = min
	}

	// Assign new value
	*c.value = cell(value)
}

// decrement decreases a cell's value.
func (c *cell) decrement() {
	// Decrement the value
	value := int64(*c.value) - 1

	// Check minimum not exceeded
	if value < *c.config.CellMin {
		value = *c.config.CellMax
	}

	// Assign new value
	*c.value = value
}

// output retrieves a cell's value.
func (c *cell) output() int64 {
	return int64(*c.value)
}

// input stores a value in a cell.
func (c *cell) input(value int64) {
	*c.value = cell(value)
}
