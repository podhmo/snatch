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
