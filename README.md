# snatch

replace struct's field value

## What is this?

Unsafe struct's field replacer. :warning: Don't use this package in production.

## Install

```console
$ go get github.com/podhmo/snatch
```

## Usecases

- debugging
- testing

## How to use

Use `Field()` or `FieldOrPanic()` method.

```go
package main

import (
	"fmt"
	"unsafe"

	"github.com/podhmo/snatch"
	"github.com/podhmo/snatch/examples/00readme/program"
)

func main() {
	prog := program.New()
	prog.Run()

	fmt.Println("----------------------------------------")
	snatch.FieldOrPanic(prog, "debug", func(ptr unsafe.Pointer) {
		realPtr := (*bool)(ptr)
		*realPtr = true
	})
	prog.Run()
}
```


Execution result.

```console
Run, debug= false
----------------------------------------
Run, debug= true
```


Then, program package's code is here.

```go
package program

import "fmt"

// Program ...
type Program struct {
	debug bool
}

func (p *Program) Run() {
	fmt.Println("Run, debug=", p.debug)
}

func New() *Program {
	return &Program{}
}
```
