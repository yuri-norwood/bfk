package bfk

type Incrementer interface {
	Increment()
}

type Decrementer interface {
	Decrement()
}

type Cell struct {
	max   int
	min   int
	Value int
}

func (cell *Cell) Increment() {
	// Increment the value
	cell.Value++

	// Check maximum not exceeded
	if cell.Value > cell.max {
		cell.Value = cell.min
	}
}

func (cell *Cell) Decrement() {
	// Decrement the value
	cell.Value--

	// Check minimum not exceeded
	if cell.Value < cell.min {
		cell.Value = cell.max
	}
}
