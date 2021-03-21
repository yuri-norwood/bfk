package test

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
		input:  nil,
		output: beer,
	}
)
