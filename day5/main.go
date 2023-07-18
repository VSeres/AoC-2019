package main

import (
	"aoc-2019/intcode"
)

func main() {
	prog := intcode.ParseFile("./input.txt")
	prog.Execute([]int{5})
}
