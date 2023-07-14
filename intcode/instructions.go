package intcode

import (
	"errors"
	"fmt"
)

type Program struct {
	memory []int
	pc     int
	modes  [3]bool
}

var (
	errInvalidAddress error = errors.New("invalid address")
)

func (p *Program) add() error {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else if len(p.memory) > opOne {
		regA = p.memory[opOne]
	} else {
		return errInvalidAddress
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else if len(p.memory) > opTwo {
		regB = p.memory[opTwo]
	} else {
		return errInvalidAddress
	}

	if len(p.memory) > dest {
		p.memory[dest] = regA + regB
		return nil
	}

	return errInvalidAddress

}

func (p *Program) multiply() error {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	var regA int
	if p.modes[2] {
		regA = opOne
	} else if len(p.memory) > opOne {
		regA = p.memory[opOne]
	} else {
		return errInvalidAddress
	}

	var regB int
	if p.modes[1] {
		regB = opTwo
	} else if len(p.memory) > opTwo {
		regB = p.memory[opTwo]
	} else {
		return errInvalidAddress
	}

	if len(p.memory) > dest {
		p.memory[dest] = regA * regB
		return nil
	}

	return errInvalidAddress
}

func (p *Program) input(input int) error {
	dest := p.memory[p.pc]
	p.pc++
	if len(p.memory) > dest {
		return errInvalidAddress
	}
	p.memory[dest] = input
	return nil
}

func (p *Program) output() error {
	dest := p.memory[p.pc]
	p.pc++
	if len(p.memory) > dest {
		return errInvalidAddress
	}
	fmt.Println("output: ", p.memory[dest])
	return nil
}
