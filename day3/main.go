package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	wires := parseFile("./example.txt")
	solve(wires)
}

type point struct {
	x int
	y int
}

type info struct {
	visited uint8
	steps   uint
}

func solve(wires [][]string) {
	pointMap := make(map[point]*info, 900)
	for wireNum, wire := range wires {
		x, y := 0, 0
		var steps uint = 0
		for _, inst := range wire {
			num, err := strconv.Atoi(inst[1:])
			if err != nil {
				log.Print(err)
			}
			switch inst[0] {
			case 'U':
				for i := y + 1; i <= y+num; i++ {
					steps += 1
					setVisited(pointMap, point{x, i}, steps, wireNum)
				}
				y += num
			case 'D':
				for i := y - 1; i >= y-num; i-- {
					steps += 1
					setVisited(pointMap, point{x, i}, steps, wireNum)
				}
				y -= num
			case 'L':
				for i := x - 1; i >= x-num; i-- {
					steps += 1
					setVisited(pointMap, point{i, y}, steps, wireNum)
				}
				x -= num
			case 'R':
				for i := x + 1; i <= x+num; i++ {
					steps += 1
					setVisited(pointMap, point{i, y}, steps, wireNum)
				}
				x += num
			}
		}
	}

	var minDist *point
	var minStep *point
	for p, pInfo := range pointMap {
		if pInfo.visited != 3 {
			continue
		}
		pCopy := p
		if minDist == nil || distance(*minDist) > distance(p) {
			minDist = &pCopy
		}
		if minStep == nil || pointMap[*minStep].steps > pInfo.steps {
			minStep = &pCopy
		}
	}
	fmt.Println("part one: ", distance(*minDist))
	fmt.Println("part two: ", pointMap[*minStep].steps)
	// print(pointMap)
}

func setVisited(visitedPoints map[point]*info, p point, steps uint, wireNum int) {
	vP := visitedPoints[p]
	if vP == nil {
		visitedPoints[p] = &info{
			visited: uint8(wireNum + 1),
			steps:   steps,
		}
	} else {
		vP.steps += steps
		vP.visited += uint8(wireNum + 1)
	}
}

func parseFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	wires := make([][]string, 0)
	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), ",")
		wires = append(wires, strArr)

	}

	return wires
}

func distance(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(n int) int {
	return ((n ^ (n >> 31)) - (n >> 31))
}

func print(pointMap map[point]info) {
	xMin, yMin, xMax, yMax := 0, 0, 0, 0
	for p := range pointMap {
		if yMax < p.y {
			yMax = p.y
		} else if yMin > p.y {
			yMin = p.y
		}
		if xMax < p.x {
			xMax = p.x
		} else if xMin > p.x {
			xMin = p.x
		}
	}

	for y := yMax; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			pointInfo := pointMap[point{x, y}]
			if pointInfo.visited == 1 || pointInfo.visited == 2 {
				fmt.Print("-")
			} else if pointInfo.visited > 1 {
				fmt.Print("X")
			} else if x == 0 && y == 0 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
