package main

import (
	"aoc-2019/intcode"
	"fmt"
)

var baseProgram intcode.Program

func main() {
	baseProgram = intcode.ParseFile("./input.txt")
	phases := make([]int, 5)
	var (
		partOneMax int
		partTwoMax int
	)
	for i := 0; i < 5*5*5*5*5; i++ {
		phases[0] += 1
		if phases[0] == 5 {
			phases[1] += 1
			phases[0] = 0
		}
		if phases[1] == 5 {
			phases[2] += 1
			phases[1] = 0
		}
		if phases[2] == 5 {
			phases[3] += 1
			phases[2] = 0
		}
		if phases[3] == 5 {
			phases[4] += 1
			phases[3] = 0
		}

		if repeatingDigit(phases) {
			continue
		}
		result := partOne(phases)
		if result > partOneMax {
			partOneMax = result
		}
		result = partTwo(phases)
		if result > partTwoMax {
			partTwoMax = result
		}
	}
	fmt.Println(partOneMax, partTwoMax)
}

func partTwo(inputPhases []int) int {
	amplifiers := make([]*intcode.Program, 5)
	for j := range amplifiers {
		prog := baseProgram.Clone()
		amplifiers[j] = &prog
	}

	phases := make([]int, 5)
	copy(phases, inputPhases)
	for i := range phases {
		phases[i] += 5
	}

	inputs := make([]int, 2)
	firstRun := true
	for !amplifiers[4].Stopped {
		for num, amplifier := range amplifiers {
			if firstRun {
				inputs[0] = phases[num]
				inputs[1] = amplifier.Execute(inputs)
			} else {
				inputs[1] = amplifier.Execute(inputs[1:])
			}
		}
		firstRun = false
	}

	return inputs[1]
}

func partOne(phases []int) int {
	amplifiers := make([]intcode.Program, 5)
	for j := range amplifiers {
		amplifiers[j] = baseProgram.Clone()
	}

	inputs := make([]int, 2)
	for num, amplifier := range amplifiers {
		inputs[0] = phases[num]
		inputs[1] = amplifier.Execute(inputs)
	}
	return inputs[1]
}

func repeatingDigit(phases []int) bool {
	for k := 0; k < 5; k++ {
		for l := k + 1; l < 5; l++ {
			if phases[k] == phases[l] {
				return true
			}
		}
	}
	return false
}
