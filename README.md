# bfk

[![Go Report Card](https://goreportcard.com/badge/github.com/yuri-norwood/bfk)](https://goreportcard.com/report/github.com/yuri-norwood/bfk)

[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter Go package

## Introduction

This package parses and executes Brainfuck code against an arbitrary `io.ReadWritter`, e.g. `os.Stdout`.

### Naming

The name [`bfk`](https://github.com/yuri-norwood/bfk) is a abriviation of *Brainfuck* (which hearafter is refered to as *BF*) in the same manner that the [`fmt`](https://golang.org/pkg/fmt/) package is a abriviation of *Format*. This seamed appropriate given the name of the language, as well as the sometimes hard-to-read names of Go package and module names.

`bfk` can be prounounced either ***bee-eff-kay*** or ***bifk***, depending on which is more comfortable to say, however, `bfk` is not an acronymn and should never be capitalised or punctuated, neither ~***BFK***~, ~***B.F.K.***~, nor ~***b.f.k.***~ may be used.

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
