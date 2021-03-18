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
