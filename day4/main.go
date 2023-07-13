package main

import (
	"fmt"
	"strconv"
)

const (
	Start = 136818
	End   = 685979
)

func main() {
	valid := 0
	for num := Start; num < End; num++ {
		if isValid(num, false) {
			valid += 1
		}
	}
	fmt.Println(valid)
}

func isValid(num int, partTwoCheck bool) bool {
	var (
		pareOneRepeat = false
		pareTwoRepeat = false
		order         = true
		numStr        = strconv.Itoa(num)
		repeatCount   = 0
	)
	for i := 0; i < len(numStr)-1; i++ {
		if numStr[i] == numStr[i+1] {
			repeatCount += 1
			pareOneRepeat = true
		} else if repeatCount == 1 {
			pareTwoRepeat = true
			repeatCount = 0
		} else {
			repeatCount = 0
		}
		if numStr[i] > numStr[i+1] {
			order = false
		}
	}
	partOne := pareOneRepeat && !partTwoCheck
	partTwo := (pareTwoRepeat || repeatCount == 1) && partTwoCheck
	return (partOne || partTwo) && order
}
