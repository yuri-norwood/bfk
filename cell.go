// This is free and unencumbered software released into the public domain.

package bfk

// cell represents a sign value within a memory tape.
type cell struct {
	value  int64
	config *Config
}

// increment increases a cell's value.
func (c *cell) increment() {
	// Get config values
	max, min := c.config.CellSize, int64(0)

	// Increment the value
	value := c.value + 1

	// Get min limit
	if c.config.CellSigned {
		min = 0 - max
	}

	// Check maximum not exceeded
	if value > max && max > 0 {
		// Check overflow
		if c.config.CellWrap {
			value = min
		} else {
			value = c.value
		}
	}

	// Assign new value
	c.value = value
}

// decrement decreases a cell's value.
func (c *cell) decrement() {
	// Get config values
	max, min := c.config.CellSize, int64(0)

	// Decrement the value
	value := c.value - 1

	// Get min limit
	if c.config.CellSigned {
		min = 0 - max
	}

	// Check minimum not exceeded
	if value < min {
		if c.config.CellWrap && max > 0 {
			// Check overflow
			value = max
		} else {
			value = c.value
		}
	}

	// Assign new value
	c.value = value
}

// output retrieves a cell's value.
func (c *cell) output() int64 {
	return c.value
}

// input stores a value in a cell.
func (c *cell) input(value int64) {
	c.value = value
}
