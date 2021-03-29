// (C) 2021 Yuri Norwood
// Use of this source code is licensed under the Unlicense,
// the terms of which can be found in the LICENSE file.

// Package bfk provides methods and types to parse, interpret,
// and compile, Brainfuck code.
//
// Parsing:
//
// 	program := bfk.Parse("+[>>>->-[>->----<<<]>>]>.---.>+..+++.>>.<.>>---.<<<.+++.------.<-.>>+.")
//
// Execution:
//
// 	var stream io.ReadWriter
// 	program.Execute(stream)
package bfk

// Parse provides a safe way of compiling a Brainfuck program
// and creating an external Program to access and execute.
func Parse(text string, settings Config) (bfk Program, err error) {
	err = nil
	bfk = &program{
		code   : text,
		config : settings,
	}

	*bfk.tape.config = &settings

	return bfk, err
}
