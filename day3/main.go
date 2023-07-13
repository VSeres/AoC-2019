package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point [2]int

func main() {
	wires := parseFile("./input.txt")
	partOne(wires)
	partTwo(wires)
}

func partOne(wires [][]string) {
	pointMap := make(map[point]uint16)

	for _, wire := range wires {
		visited := make(map[point]bool)
		x, y := 0, 0
		for _, inst := range wire {
			num, err := strconv.Atoi(inst[1:])
			if err != nil {
				log.Print(err)
			}
			if inst[0] == 'U' {
				for i := y + 1; i <= y+num; i++ {
					visited[point{x, i}] = true
				}
				y += num
			} else if inst[0] == 'D' {
				for i := y - 1; i >= y-num; i-- {
					visited[point{x, i}] = true
				}
				y -= num
			} else if inst[0] == 'L' {
				for i := x - 1; i >= x-num; i-- {
					visited[point{i, y}] = true
				}
				x -= num
			} else if inst[0] == 'R' {
				for i := x + 1; i <= x+num; i++ {
					visited[point{i, y}] = true
				}
				x += num
			}
		}
		for p := range visited {
			pointMap[p] += 1
		}
	}

	var min point
	isSet := false
	for p, value := range pointMap {
		if value < 2 {
			continue
		}
		if !isSet || distance(min) > distance(p) {
			min = p
			isSet = true
		}
	}
	log.Print(min, distance(min))
	// print(pointMap)
}

func partTwo(wires [][]string) {
	type info [2]uint
	pointMap := make(map[point]info)

	for _, wire := range wires {
		visited := make(map[point]uint)
		x, y := 0, 0
		var step uint = 0

		for _, inst := range wire {

			num, _ := strconv.Atoi(inst[1:])
			if inst[0] == 'U' {
				for i := y + 1; i <= y+num; i++ {
					step += 1
					visited[point{x, i}] = step
				}
				y += num
			} else if inst[0] == 'D' {
				for i := y - 1; i >= y-num; i-- {
					step += 1
					visited[point{x, i}] = step
				}
				y -= num
			} else if inst[0] == 'L' {
				for i := x - 1; i >= x-num; i-- {
					step += 1
					visited[point{i, y}] = step
				}
				x -= num
			} else if inst[0] == 'R' {
				for i := x + 1; i <= x+num; i++ {
					step += 1
					visited[point{i, y}] = step
				}
				x += num
			}

		}
		for p, step := range visited {
			pInfo := pointMap[p]
			if pInfo[0] != 0 {
				pInfo[0] += 1
				pInfo[1] += step
				pointMap[p] = pInfo
			} else {
				pointMap[p] = info{1, step}
			}
		}
	}

	var min point
	isSet := false
	for p, value := range pointMap {
		if value[0] < 2 {
			continue
		}
		if !isSet || pointMap[min][1] > value[1] {
			min = p
			isSet = true
		}
	}
	log.Print(min, pointMap[min][1])
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

func distance(p Point) int {
	return abs(p.x) + abs(p.y)
}

func abs(n int) int {
	return ((n ^ (n >> 31)) - (n >> 31))
}

func print(pointMap map[point]uint16) {
	xMin, yMin, xMax, yMax := 0, 0, 0, 0
	for p := range pointMap {
		if yMax < p[1] {
			yMax = p[1]
		} else if yMin > p[1] {
			yMin = p[1]
		}
		if xMax < p[0] {
			xMax = p[0]
		} else if xMin > p[0] {
			xMin = p[0]
		}
	}

	for y := yMax; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			if pointMap[point{x, y}] == 1 {
				fmt.Print("-")
			} else if pointMap[point{x, y}] > 1 {
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
