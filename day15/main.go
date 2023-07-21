package main

import (
	"aoc-2019/intcode"
	"fmt"
	"time"
)

type point struct {
	x, y int
}

func (p point) dir(other point) int {
	if p.x > other.x {
		return left
	}
	if p.x < other.x {
		return right
	}
	if p.y > other.y {
		return down
	}
	if p.y < other.y {
		return up
	}
	return -1
}

var roomMap = make(map[point]int, 0)
var robotPos = point{}
var nodeStack = make([]point, 0)

const (
	empty = 1
	goal  = 2
	wall  = 3
)

const (
	up    = 1
	down  = 2
	left  = 3
	right = 4
)

var dir = left
var program intcode.Program

func main() {
	program = intcode.ParseFile("input.txt")
	roomMap[robotPos] = empty
	nodeStack = append(nodeStack, robotPos)
	stepC := 0
	for stepC < 10 {
		lookAround()
		newPos := getNewPos(dir)
		if roomMap[newPos] == wall {
			nodeStack = append(nodeStack, robotPos)
			for {
				newDir := ((dir + 1) % 5)
				if newDir == 0 {
					newDir++
				}
				dir = newDir
				newPos = getNewPos(dir)
				if nodeStack[len(nodeStack)-1].dir(newPos) == dir && roomMap[newPos] == wall {
					break
				}
			}
		}

		move(dir)
		lookAround()
		stepC++
		display()
		time.Sleep(1 * time.Second)
	}

}

func lookAround() {
	upP, rightP, downP, leftP := getDirCoords()

	if roomMap[upP] == 0 && move(up) {
		move(down)
	}
	if roomMap[rightP] == 0 && move(right) {
		move(left)
	}
	if roomMap[downP] == 0 && move(down) {
		move(up)
	}
	if roomMap[leftP] == 0 && move(left) {
		move(right)
	}
}

func display() {
	min := point{}
	max := point{}
	for pos := range roomMap {
		if min.x > pos.x {
			min.x = pos.x
		}
		if min.y > pos.y {
			min.y = pos.y
		}
		if max.x < pos.x {
			max.x = pos.x
		}
		if max.y < pos.y {
			max.y = pos.y
		}
	}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			p := roomMap[point{x, y}]
			if robotPos.x == x && robotPos.y == y && p != goal {
				fmt.Print("r")
				continue
			}
			switch p {
			case 0:
				fmt.Print("?")
			case empty:
				fmt.Print(".")
			case wall:
				fmt.Print("#")
			case goal:
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func move(dir int) bool {
	program.SetInput(dir)
	newPos := getNewPos(dir)

	output := program.Execute()
	updateMap(output[0], newPos)
	if output[0] != 0 {
		robotPos = newPos
		return true
	}
	return false
}

func updateMap(tile int, newPos point) {
	if tile == 0 {
		roomMap[newPos] = wall
	} else if tile == 1 {
		roomMap[newPos] = empty
	} else {
		roomMap[newPos] = goal
	}
}

func getNewPos(dir int) point {
	switch dir {
	case up:
		return point{robotPos.x, robotPos.y + 1}
	case right:
		return point{robotPos.x + 1, robotPos.y}
	case down:
		return point{robotPos.x, robotPos.y - 1}
	case left:
		return point{robotPos.x - 1, robotPos.y}
	}
	return point{}
}

func backAndForth(dir int) {
	input := make([]int, 2)
	input[0] = dir
	if dir == up {
		input[1] = down
	}
	if dir == down {
		input[1] = up
	}
	if dir == right {
		input[1] = left
	}
	if dir == left {
		input[1] = right
	}
	newPos := getNewPos(dir)
	program.SetInputs(input)
	out := program.Execute()
	updateMap(out[0], newPos)
	fmt.Println(out, dir)
}

func getDirCoords() (point, point, point, point) {
	up := point{robotPos.x, robotPos.y + 1}
	right := point{robotPos.x + 1, robotPos.y}
	down := point{robotPos.x, robotPos.y - 1}
	left := point{robotPos.x - 1, robotPos.y}

	return up, right, down, left
}
