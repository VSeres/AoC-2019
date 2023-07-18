package intcode

type Program struct {
	memory       []int
	pc           int
	modes        [3]byte
	inputArr     []int
	inputPointer int
	halt         bool
	Stopped      bool
	base         int
}

func (p *Program) add() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	regA := p.getValue(p.modes[2], opOne)
	regB := p.getValue(p.modes[1], opTwo)

	p.putValue(dest, regA+regB, p.modes[0])

}

func (p *Program) multiply() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	regA := p.getValue(p.modes[2], opOne)
	regB := p.getValue(p.modes[1], opTwo)

	p.putValue(dest, regA*regB, p.modes[0])

}

func (p *Program) input() {
	input := p.inputArr[p.inputPointer]
	p.inputPointer += 1

	dest := p.memory[p.pc]
	p.pc++
	p.putValue(dest, input, p.modes[2])
}

func (p *Program) output() int {
	opOne := p.memory[p.pc]
	p.pc++
	return p.getValue(p.modes[2], opOne)

}

func (p *Program) jumpIf(nonZero bool) {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	p.pc += 2

	regA := p.getValue(p.modes[2], opOne)
	regB := p.getValue(p.modes[1], opTwo)

	if (regA != 0) == nonZero {
		p.pc = regB
	}
}

func (p *Program) lessThan() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	regA := p.getValue(p.modes[2], opOne)
	regB := p.getValue(p.modes[1], opTwo)

	if regA < regB {
		p.putValue(dest, 1, p.modes[0])
	} else {
		p.putValue(dest, 0, p.modes[0])
	}
}

func (p *Program) equals() {
	opOne := p.memory[p.pc]
	opTwo := p.memory[p.pc+1]
	dest := p.memory[p.pc+2]
	p.pc += 3

	regA := p.getValue(p.modes[2], opOne)
	regB := p.getValue(p.modes[1], opTwo)

	if regA == regB {
		p.putValue(dest, 1, p.modes[0])
	} else {
		p.putValue(dest, 0, p.modes[0])
	}
}

func (p *Program) modBase() {
	param := p.memory[p.pc]
	p.pc++
	p.base += p.getValue(p.modes[2], param)
}
