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

func partOne(baseCode []int) {
	code := make([]int, len(baseCode))
	copy(code, baseCode)
	log.Print(intcode.Execute(code))
}

func partTwo(baseCode []int) {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			code := make([]int, len(baseCode))
			copy(code, baseCode)
			(code)[1] = i
			(code)[2] = j
			result := intcode.Execute(code)
			if result == 19690720 {
				log.Print(result, i, j)
				return
			} else if result > 19690720 {
				break
			}
		}
	}
}
