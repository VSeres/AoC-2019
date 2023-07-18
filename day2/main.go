package main

import (
	"aoc-2019/intcode"
	"log"
)

func main() {
	baseCode := intcode.ParseFile("./input.txt")
	partOne(baseCode)
	partTwo(baseCode)

}

func partOne(baseProgram intcode.Program) {
	program := baseProgram.Clone()
	program.WriteMemory(1, 12)
	program.WriteMemory(2, 2)
	program.Execute([]int{})
	log.Print(program.ReadMemory(0))
}

func partTwo(baseProgram intcode.Program) {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			program := baseProgram.Clone()
			program.WriteMemory(1, i)
			program.WriteMemory(2, j)
			program.Execute([]int{})
			result := program.ReadMemory(0)
			if result == 19690720 {
				log.Print(result, i, j)
				return
			} else if result > 19690720 {
				break
			}
		}
	}
}
