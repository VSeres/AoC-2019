package main

import (
	"aoc-2019/intcode"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var scaffold = make(map[point]bool, 0)
var robot point
var dir int

func main() {
	program := intcode.ParseFile("input.txt")
	output := program.Execute()
	steps := getSteps(output)
	vacum(program, steps)
}

func vacum(program intcode.Program, steps []string) {
	inst := patterSearch()
	program.WriteMemory(0, 2)
	input := make([]int, 0)

	routStr := getRoutString()
	for i, patt := range inst {
		reg := regexp.MustCompile("(" + patt + ")")
		routStr = reg.ReplaceAllString(routStr, string(rune(65+i)))
	}
	routStr = strings.ReplaceAll(routStr, " ", ",")
	for _, v := range routStr {
		input = append(input, int(v))
	}
	input = append(input, 10)

	for _, inst := range inst {
		function := strings.ReplaceAll(inst, " ", ",")
		for _, v := range function {
			input = append(input, int(v))
		}
		input = append(input, 10)
	}
	input = append(input, 'n')
	input = append(input, 10)
	// for _, v := range input {
	// 	fmt.Printf("%c", v)
	// }
	program.SetInputs(input)
	out := program.Execute()
	fmt.Println(out[len(out)-1])
}

func getSteps(output []int) []string {
	var pos point
	for _, v := range output {
		switch v {
		case '^':
			robot = pos
			dir = 0
		case '>':
			robot = pos
			dir = 1
		case 'v':
			robot = pos
			dir = 2
		case '<':
			robot = pos
			dir = 3
		case 35:
			scaffold[pos] = true
		}

		if v == 10 {
			pos.y += 1
			pos.x = 0
		} else {
			pos.x += 1
		}
	}
	alignment := 0
	inter := 0
	for p := range scaffold {
		neighbours := neighbourCount(p)
		if neighbours > 3 {
			inter += 1
			alignment += p.x * p.y
		}
	}
	travers()
	return route
}

func neighbourCount(p point) int {
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
	return neighbours
}

var route = make([]string, 0)
var visited = make(map[point]bool, len(scaffold))

func patterSearch() []string {
	inst := make([]string, 3)
	routeString := getRoutString()
	whitespace := regexp.MustCompile(`\s{2,}`)
	for i := range inst {
		var pattern string
		var patternBuff strings.Builder
		matchCount := 0
		for i := 0; i < len(routeString)-1; i += 2 {
			patternBuff.WriteByte(routeString[i])
			patternBuff.WriteByte(routeString[i+1])
			if patternBuff.Len() > 20 {
				break
			}
			reg := regexp.MustCompile(patternBuff.String())
			count := len(reg.FindAllString(routeString, -1)) * patternBuff.Len()
			if (count > 1 && count > matchCount) || patternBuff.Len() <= 6 {
				pattern = patternBuff.String()
				matchCount = count
			}
		}
		trimmed := strings.Trim(pattern, " ")
		reg := regexp.MustCompile(pattern)
		routeString = reg.ReplaceAllString(routeString, "")
		routeString = whitespace.ReplaceAllString(strings.Trim(routeString, " "), " ")
		inst[i] = trimmed

	}
	return inst
}

func getRoutString() string {
	var buff strings.Builder
	for _, v := range route {
		if buff.Len() > 0 {
			buff.WriteByte(' ')
		}
		buff.WriteString(v)
	}
	return buff.String()
}

func travers() {
	done := false
	fmt.Print("\033[s")
	for !done {
		dist := 0
		setNewDir()
		for hitWall() {
			dist += 1
			robot = getNewPos(dir)
			visited[robot] = true
		}
		num := strconv.Itoa(dist)
		route = append(route, num)

		display()
		// time.Sleep(1 * time.Second)
		done = neighbourCount(robot) <= 1
	}
}

func hitWall() bool {
	return (dir == 0 && scaffold[robot.down()]) ||
		(dir == 1 && scaffold[robot.right()]) ||
		(dir == 2 && scaffold[robot.up()]) ||
		(dir == 3 && scaffold[robot.left()])
}

func setNewDir() {
	newDir := -1
	switch dir {
	case 0:
		if scaffold[robot.right()] {
			newDir = 1
			route = append(route, "R")
		} else if scaffold[robot.left()] {
			newDir = 3
			route = append(route, "L")
		}
	case 1:
		if scaffold[robot.down()] {
			newDir = 0
			route = append(route, "L")
		} else if scaffold[robot.up()] {
			newDir = 2
			route = append(route, "R")
		}
	case 2:
		if scaffold[robot.right()] {
			newDir = 1
			route = append(route, "L")
		} else if scaffold[robot.left()] {
			newDir = 3
			route = append(route, "R")
		}
	case 3:
		if scaffold[robot.down()] {
			newDir = 0
			route = append(route, "R")
		} else if scaffold[robot.up()] {
			newDir = 2
			route = append(route, "L")
		}
	}
	dir = newDir
}

func getNewPos(dir int) point {
	switch dir {
	case 0:
		return robot.down()
	case 1:
		return robot.right()
	case 2:
		return robot.up()
	case 3:
		return robot.left()
	}
	return point{}
}

func display() {
	var max point
	for k := range scaffold {
		if max.x < k.x {
			max.x = k.x
		}
		if max.y < k.y {
			max.y = k.y
		}
	}
	fmt.Println("\033[u")
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			p := point{x, y}
			if scaffold[p] {
				if visited[p] {
					fmt.Print("\033[1;32m#\033[0m")
				} else {
					fmt.Print("\033[1;31m#\033[0m")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
