package intcode

import "fmt"

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
	} else {
		regA = p.memory[opOne]
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else {
		regB = p.memory[opTwo]
	}
	// fmt.Println("add ", p.memory[p.pc-4], regA, regB)
	p.memory[dest] = regA + regB
}

func (p *Program) multiply() {
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

	p.memory[dest] = regA * regB
}

func (p *Program) input(input int) {
	dest := p.memory[p.pc]
	p.pc++

	p.memory[dest] = input
}

func (p *Program) output() {
	dest := p.memory[p.pc]
	p.pc++

	fmt.Println("output: ", p.memory[dest])
}
