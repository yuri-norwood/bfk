// This is free and unencumbered software released into the public domain.

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
func Parse(text string) (bfk Program, err error) {
	err = nil
	bfk = &program{code: text}

	return bfk, err
}
