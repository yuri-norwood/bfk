# bfk

[![Go Report Card](https://goreportcard.com/badge/github.com/yuri-norwood/bfk)](https://goreportcard.com/report/github.com/yuri-norwood/bfk)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuri-norwood/bfk.svg)](https://pkg.go.dev/github.com/yuri-norwood/bfk)

Brainfuck interpreter Go package

## Introduction

This package parses and executes Brainfuck code against an arbitrary
`io.ReadWritter` implementation, e.g. `os.Stdout`.

### Naming

The name [`bfk`](https://github.com/yuri-norwood/bfk) is a abriviation of
*[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck)* (which hearafter is
refered to as *BF*) in the same manner that the
[`fmt`](https://golang.org/pkg/fmt/) package is a abriviation of *Format*.
This seamed appropriate given profane the name of the language, as well as
the sometimes hard-to-read names of Go package and module names.

`bfk` can be prounounced either ***bee-eff-kay*** or ***bifk***, depending on
which is more comfortable to say, however, `bfk` is not an acronymn and should
never be capitalised or punctuated, neither ~***BFK***~, ~***B.F.K.***~, nor
~***b.f.k.***~ may be used.

## Example

```go
package main

import (
    "os"

    "github.com/yuri-norwood/bfk"
)

const HelloWorld = "+[>>>->-[>->----<<<]>>]>.---.>+..+++.>>.<.>>---.<<<.+++.------.<-.>>+."

func main() {
    program := bfk.Parse(HelloWorld)
    program.Execute(os.Stdout) // Prints "hello, world!"
}

```

## Usage

`bfk` allows any implementation of the `io.ReadWritter` interface to be used by
*BF* code, this is not limited to simple Hello World type programs that simply
write a string to standard output, `bfk` allows the *BF* code to do *anything*,
such as writing to images, preforming byte-wise encoding and decoding, and even
modifying sound streams. In addition to this, parsing the *BF* source code, and
execution of the source code, are separated, meaning that the same `bfk.Program`
can be reused on multiple `io.ReadWritter` sources, or even allowing for
parralelization by running many copies in their own go routines. Note however,
that `bfk` leaves concurency up to the user, as *BF* itself is inhearently
single-threaded.
