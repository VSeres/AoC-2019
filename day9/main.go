package main

import (
	"aoc-2019/intcode"
	"fmt"
)

func main() {
	program := intcode.ParseFile("input.txt")
	fmt.Println(program.Execute([]int{2}))
}
