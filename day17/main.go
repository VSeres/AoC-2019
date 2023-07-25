package main

import (
	"aoc-2019/intcode"
	"fmt"
	"strings"
)

var scaffold = make(map[point]bool, 0)
var robot point

func main() {
	program := intcode.ParseFile("input.txt")
	output := program.Execute()
	var buff strings.Builder
	var pos point
	for _, v := range output {
		if v == 35 {
			scaffold[pos] = true
		}
		if v == 10 {
			pos.y += 1
			pos.x = 0
		} else {
			pos.x += 1
		}
		buff.WriteByte(byte(v))
	}
	fmt.Println(buff.String())
	var max point
	for k := range scaffold {
		if max.x < k.x {
			max.x = k.x
		}
		if max.y < k.y {
			max.y = k.y
		}
	}
	for y := 0; y < max.y; y++ {
		for x := 0; x < max.x; x++ {
			if scaffold[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	alignment := 0
	inter := 0
	for p := range scaffold {
		neighbours := 0
		if scaffold[p.up()] {
			neighbours++
		}
		if scaffold[p.down()] {
			neighbours++
		}
		if scaffold[p.left()] {
			neighbours++
		}
		if scaffold[p.right()] {
			neighbours++
		}
		if neighbours > 3 {
			inter += 1
			alignment += p.x * p.y
			fmt.Println(p)
		}
	}
	fmt.Println(alignment, inter)
}
