package intcode

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func (p *Program) Execute() {
	for p.pc < len(p.memory) {
		inst := p.memory[p.pc]
		p.pc++
		opCode := inst % 100
		if opCode == 99 {
			break
		}
		modesStr := fmt.Sprintf("%05d", inst)

		modesStr = modesStr[:len(modesStr)-2]
		p.modes = [3]bool{}
		for i := len(modesStr) - 1; i >= 0; i-- {
			p.modes[i] = modesStr[i] == '1'
		}

		switch opCode {
		case 1:
			p.add()
		case 2:
			p.multiply()
		case 3:
			p.input(1)
		case 4:
			p.output()
		}
	}
	// fmt.Println(p.memory)
}

func ParseFile(path string) Program {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	buff := make([]byte, 2048)
	n, err := file.Read(buff)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	buff = buff[:n]
	str := strings.Trim(string(buff), "\n\r")
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
