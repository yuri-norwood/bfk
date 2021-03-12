package bfk

var (
	cell_max int64 = 255
	cell_min int64 = 0
)

type Incrementer interface {
	Increment()
}

type Decrementer interface {
	Decrement()
}

type IncrementDecrementer interface {
	Decrementer
	Incrementer
}

type Cell int64

func (cell *Cell) Increment() {
	// Increment the value
	value := int64(*cell) + 1

	// Check maximum not exceeded
	if value > cell_max {
		value = cell_min
	}

	// Asign new value
	*cell = Cell(value)
}

func (cell *Cell) Decrement() {
	// Decrement the value
	value := int64(*cell) - 1

	// Check minimum not exceeded
	if value < cell_min {
		value = cell_max
	}

	// Asign new value
	*cell = Cell(value)
}
