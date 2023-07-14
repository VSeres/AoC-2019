package intcode

import (
	"fmt"
)

type Program struct {
	memory []int
	pc     int
	modes  [3]bool
}

func (p *Program) add() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else if len(p.memory) > opOne {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else if len(p.memory) > opTwo {
		regB = p.memory[opTwo]
	}
	if len(p.memory) > dest {
		p.memory[dest] = regA + regB
	}
}

func (p *Program) multiply() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else if len(p.memory) > opOne {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else if len(p.memory) > opTwo {
		regB = p.memory[opTwo]
	}

	if len(p.memory) > dest {
		p.memory[dest] = regA * regB
	}
}

func (p *Program) input(input int) {
	dest := p.memory[p.pc]
	p.pc++
	p.memory[dest] = input
}

func (p *Program) output() {
	opOne := p.memory[p.pc]
	p.pc++
	if p.modes[2] {
		fmt.Println("output: ", opOne)
	} else {
		fmt.Println("output: ", p.memory[opOne])
	}
}

func (p *Program) jumpIf(nonZero bool) {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	p.pc += 2
	var regA int
	if p.modes[2] {
		regA = opOne
	} else {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else {
		regB = p.memory[opTwo]
	}

	if (regA != 0) == nonZero {
		p.pc = regB
	}
}

func (p *Program) lessThan() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else {
		regB = p.memory[opTwo]
	}

	if regA < regB {
		p.memory[dest] = 1
	} else {
		p.memory[dest] = 0
	}
}

func (p *Program) equals() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else {
		regB = p.memory[opTwo]
	}

	if regA == regB {
		p.memory[dest] = 1
	} else {
		p.memory[dest] = 0
	}
}
