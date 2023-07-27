package main

import (
	"fmt"
	"strconv"
)

func solveOld(wires [][]string) {
	pointMap := make(map[point]info)
	for _, wire := range wires {
		visited := make(map[point]*info)
		x, y := 0, 0
		var steps uint = 0
		for _, inst := range wire {
			num, err := strconv.Atoi(inst[1:])
			if err != nil {
				fmt.Println(err)
			}
			if inst[0] == 'U' {
				for i := y + 1; i <= y+num; i++ {
					steps += 1
					setVisitedOld(visited, point{x, i}, steps)
				}
				y += num
			} else if inst[0] == 'D' {
				for i := y - 1; i >= y-num; i-- {
					steps += 1
					setVisitedOld(visited, point{x, i}, steps)
				}
				y -= num
			} else if inst[0] == 'L' {
				for i := x - 1; i >= x-num; i-- {
					steps += 1
					setVisitedOld(visited, point{i, y}, steps)
				}
				x -= num
			} else if inst[0] == 'R' {
				for i := x + 1; i <= x+num; i++ {
					steps += 1
					setVisitedOld(visited, point{i, y}, steps)
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
	// fmt.Println("part one: ", distance(*minDist))
	// fmt.Println("part two: ", pointMap[*minStep].steps)
	// print(pointMap)
}

func setVisitedOld(visited map[point]*info, p point, steps uint) {
	if visited[p] == nil {
		visited[p] = &info{
			visited: 1,
			steps:   steps,
		}
	}
}
