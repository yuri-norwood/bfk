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

type Lefter interface {
	Left()
}

type Righter interface {
	Right()
}

type LeftRighter interface {
	Lefter
	Righter
}

type Outputter interface {
	Output() int64
}

type Inputter interface {
	Input(int64)
}

type OutputInputter interface {
	Outputter
	Inputter
}
