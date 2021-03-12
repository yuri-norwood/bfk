package bfk

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
