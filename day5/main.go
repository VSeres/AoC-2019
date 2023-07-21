package main

import (
	"aoc-2019/intcode"
)

func main() {
	prog := intcode.ParseFile("./input.txt")
	prog.SetInput(5)
	prog.Execute()
}
