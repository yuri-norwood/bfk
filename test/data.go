package test

import _ "embed" // allow embedding test files directly into code

type code struct {
	name   string
	source string
	input  []byte
	output []byte
}

var (
	//go:embed _data/rdebath/testing/Beer.b
	beerSource string

	//go:embed _data/rdebath/testing/Beer.out
	beerOutput []byte

	beer = code{
		name:   "99 Bottles of Beer on the Wall",
		source: beerSource,
		output: beerOutput,
	}
)

var (
	//go:embed _data/rdebath/bitwidth.b
	bitwidthSource string

	bitwidth = code{
		name:   "Cell Bit Width Detection (",
		source: bitwidthSource,
		output: []byte("Hello World! 255"), // output varies, this is the target
	}
)

var (
	//go:embed _data/rdebath/testing/Collatz.b
	collatzSource string

	//go:embed _data/rdebath/testing/Collatz.in
	collatzInput []byte

	//go:embed _data/rdebath/testing/Collatz.out
	collatzOutput []byte

	// collatz is a program that reads numbers, one per line, and outputs
	// the length of sequence, starting at the number, and following the
	// collatz sequence down to 1. Note, there appears to be no upper limit
	// on the size of the input numbers, the provided example is the integer
	// representing the first 2000 digits of PI. (Runs unexpectedly fast).
	collatz = code{
		name:   "Collatz Starting from PI x 2000",
		source: collatzSource,
		input:  collatzInput,
		output: collatzOutput,
	}
)
