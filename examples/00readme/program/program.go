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
