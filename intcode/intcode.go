package intcode

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (p *Program) Execute() (output []int) {
	for p.pc < len(p.memory) {
		opcode := p.readInsctrucion()
		switch opcode {
		case 1:
			p.add()
		case 2:
			p.multiply()
		case 3:
			if p.inputPointer >= len(p.inputArr) {
				p.halt = true
				return
			}
			p.input()
		case 4:
			output = append(output, p.output())
		case 5:
			p.jumpIf(true)
		case 6:
			p.jumpIf(false)
		case 7:
			p.lessThan()
		case 8:
			p.equals()
		case 9:
			p.modBase()
		case 99:
			p.Stopped = true
			return
		default:
			fmt.Printf("invaild opcode: %d\n", opcode)
		}
	}
	return output
}

func (p *Program) SetInput(i int) {
	if len(p.inputArr) > 1 || len(p.inputArr) == 0 {
		p.inputArr = make([]int, 1)
	}
	p.inputPointer = 0
	p.inputArr[0] = i
}

func (p *Program) SetInputs(inputs []int) {
	p.inputArr = inputs
	p.inputPointer = 0
}

func (p *Program) readInsctrucion() int {
	if p.halt {
		p.pc -= 1
		p.halt = false
	}
	inst := p.memory[p.pc]
	p.pc++
	opcode := inst % 100
	if opcode == 99 {
		return 99
	}
	modesStr := fmt.Sprintf("%05d", inst)
	modesStr = modesStr[:len(modesStr)-2]
	for i := len(modesStr) - 1; i >= 0; i-- {
		p.modes[i] = modesStr[i]
	}

	return opcode
}

func ParseFile(path string) Program {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var builder strings.Builder
	buff := make([]byte, 2048)
	var n int

	for err != io.EOF {
		n, err = file.Read(buff)

		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}
		builder.Write(buff[:n])
	}

	str := strings.Trim(builder.String(), "\n\r")
	strArr := strings.Split(str, ",")
	code := make([]int, len(strArr))

	for i := range code {
		if strings.Trim(strArr[i], " \n\r") == "" {
			continue
		}
		num, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		code[i] = num
	}

	return Program{memory: code}
}

func (p *Program) Clone() Program {
	program := Program{
		memory: make([]int, len(p.memory)),
	}
	copy(program.memory, p.memory)

	return program
}

func (p Program) ReadMemory(address int) int {
	if address >= len(p.memory) {
		return -1
	}

	return p.memory[address]
}

func (p Program) WriteMemory(address int, value int) {
	if address >= len(p.memory) {
		return
	}
	p.memory[address] = value
}
