package intcode

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute(code []int) int {
	codeLen := len(code)
	for pc := 0; pc < len(code); pc += 4 {
		inst := code[pc]
		if inst == 99 {
			break
		}
		op1Addr := code[pc+1]
		op2Addr := code[pc+2]
		destinasionAddr := code[pc+3]
		if codeLen <= destinasionAddr || codeLen <= op1Addr || codeLen <= op2Addr {
			return -10
		}
		switch inst {
		case 1:
			code[destinasionAddr] = code[op1Addr] + code[op2Addr]
		case 2:
			code[destinasionAddr] = code[op1Addr] * code[op2Addr]
		}

	}
	return code[0]
}

func ParseFile(path string) []int {
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
		num, err := strconv.Atoi(strArr[i])
		if err != nil {
			log.Print(err)
			continue
		}
		code[i] = num
	}

	return code
}
