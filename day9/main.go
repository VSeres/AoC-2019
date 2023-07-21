package main

import (
	"aoc-2019/intcode"
	"fmt"
)

func main() {
	program := intcode.ParseFile("input.txt")
	program.SetInputs([]int{2})
	fmt.Println(program.Execute())
}
