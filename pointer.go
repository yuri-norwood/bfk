package bfk

var (
	pointer_max int64 = 29000
	pointer_min int64 = 0
)

type Pointer int64

func (pointer *Pointer) Increment() {
	// Increment the value
	value := int64(*pointer) + 1

	// Check maximum not exceeded
	if value > pointer_max {
		value = pointer_min
	}

	// Asign new value
	*pointer = Pointer(value)
}

func (pointer *Pointer) Decrement() {
	// Decrement the value
	value := int64(*pointer) - 1

	// Check minimum not exceeded
	if value < pointer_min {
		value = pointer_max
	}

	// Asign new value
	*pointer = Pointer(value)
}
