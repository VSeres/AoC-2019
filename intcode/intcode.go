package intcode

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func (p *Program) Execute(inputs []int) (output int) {
	p.inputArr = inputs
	p.inputPointer = 0
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
			output = p.output()
		case 5:
			p.jumpIf(true)
		case 6:
			p.jumpIf(false)
		case 7:
			p.lessThan()
		case 8:
			p.equals()
		case 99:
			p.Stopped = true
			return
		default:
			log.Printf("invaild opcode: %d", opcode)
		}
	}
	return
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
	p.modes = [3]bool{}
	for i := len(modesStr) - 1; i >= 0; i-- {
		p.modes[i] = modesStr[i] == '1'
	}
	return opcode
}

func ParseFile(path string) Program {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var builder strings.Builder

	buff := make([]byte, 2048)
	var n int
	for err != io.EOF {
		n, err = file.Read(buff)

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		builder.Write(buff[:n])
	}

	str := strings.Trim(builder.String(), "\n\r")
	strArr := strings.Split(str, ",")
	code := make([]int, len(strArr))

	for i := range code {
		fmt.Printf("\r\033[K%d/%d", i+1, len(strArr))
		if strings.Trim(strArr[i], " \n\r") == "" {
			continue
		}
		num, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Println()
			log.Fatal(err)
		}
		code[i] = num
	}
	fmt.Println()

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
