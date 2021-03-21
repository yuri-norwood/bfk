package test

import _ "embed"

type code struct {
	name   string
	source string
	input  []byte
	output []byte
}

//go:embed _data/rdebath/testing/

var (
	//go:embed _data/rdebath/testing/Beer.b
	beerSource string

	//go:embed _data/rdebath/testing/Beer.out
	beerOutput []byte

	beer = code{
		name:   "99 Bottles of Beer on the Wall",
		source: beer,
		output: beer,
	}
)

var (
	//go:embed _data_/rdebath/bitwidth.b
	bitwidthSource string

	bitwidth = code{
		name:   "Cell Bit Width Detection (",
		source: bitwidthSource,
		output: "Hello World! 255", // output varies, this is the target
	}
)
