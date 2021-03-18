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
