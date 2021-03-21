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

	collatz = code{
		name:   "Collatz Starting from PI x 1 bazillion",
		source: collatzSource,
		input:  collatzInput,
		output: collatzOutput,
	}
)
