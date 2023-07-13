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
	pointMap := make(map[point]info)
	for _, wire := range wires {
		visited := make(map[point]*info)
		x, y := 0, 0
		var steps uint = 0
		for _, inst := range wire {
			num, err := strconv.Atoi(inst[1:])
			if err != nil {
				log.Print(err)
			}
			if inst[0] == 'U' {
				for i := y + 1; i <= y+num; i++ {
					steps += 1
					setVisited(visited, point{x, i}, steps)
				}
				y += num
			} else if inst[0] == 'D' {
				for i := y - 1; i >= y-num; i-- {
					steps += 1
					setVisited(visited, point{x, i}, steps)
				}
				y -= num
			} else if inst[0] == 'L' {
				for i := x - 1; i >= x-num; i-- {
					steps += 1
					setVisited(visited, point{i, y}, steps)
				}
				x -= num
			} else if inst[0] == 'R' {
				for i := x + 1; i <= x+num; i++ {
					steps += 1
					setVisited(visited, point{i, y}, steps)
				}
				x += num
			}
		}

		for p, vInfo := range visited {
			pInfo := pointMap[p]
			pInfo.steps += vInfo.steps
			pInfo.visited += 1
			pointMap[p] = pInfo
		}
	}
	var minDist *point
	var minStep *point
	for p, pInfo := range pointMap {
		if pInfo.visited < 2 {
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
	log.Print("part one: ", distance(*minDist))
	log.Print("part two: ", pointMap[*minStep].steps)
	// print(pointMap)
}

func setVisited(visited map[point]*info, p point, steps uint) {
	if visited[p] == nil {
		visited[p] = &info{
			visited: 1,
			steps:   steps,
		}
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
			if pointInfo.visited == 1 {
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
