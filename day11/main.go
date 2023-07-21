package main

import (
	"aoc-2019/intcode"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type point struct {
	x int
	y int
}

type robot struct {
	point
	facing int
}

var paintingRobot = robot{}

func main() {
	code := intcode.ParseFile("input.txt")
	partOne(code.Clone())
	partTwo(code.Clone())
}

func partOne(code intcode.Program) {
	panelMap := make(map[point]int, 0)
	panelMap[point{0, 0}] = 1
	paint(code, panelMap)
	fmt.Println(len(panelMap))
}

func partTwo(code intcode.Program) {
	paintingRobot = robot{}
	panelMap := make(map[point]int, 0)
	panelMap[point{0, 0}] = 1
	paint(code, panelMap)
	minX := 500
	minY := 500
	maxX := 0
	maxY := 0
	for p := range panelMap {
		if minX > p.x {
			minX = p.x
		}
		if minY > p.y {
			minY = p.y
		}
		if maxX < p.x {
			maxX = p.x
		}
		if maxY < p.y {
			maxY = p.y
		}
	}

	img := image.NewGray(image.Rect(0, 0, maxX-minX+1, maxY-minY+1))
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if panelMap[point{x: x, y: y}] == 1 {
				img.SetGray(x-minX, maxY-y, color.Gray{255})
			} else {
				img.SetGray(x-minX, maxY-y, color.Gray{0})
			}
		}
	}
	file, _ := os.Create("./part_two.png")
	png.Encode(file, img)

}

func paint(code intcode.Program, panelMap map[point]int) {
	for !code.Stopped {
		color := panelMap[paintingRobot.point]
		code.SetInput(color)
		output := code.Execute()

		panelMap[paintingRobot.point] = output[0]
		y := 0
		x := 0
		if (output[1] == 0 && paintingRobot.facing == 0) || (output[1] == 1 && paintingRobot.facing == 2) {
			x = -1
		} else if (output[1] == 1 && paintingRobot.facing == 0) || (output[1] == 0 && paintingRobot.facing == 2) {
			x = 1
		}
		if (output[1] == 1 && paintingRobot.facing == 1) || (output[1] == 0 && paintingRobot.facing == 3) {
			y = -1
		} else if (output[1] == 0 && paintingRobot.facing == 1) || (output[1] == 1 && paintingRobot.facing == 3) {
			y = 1
		}

		if x > 0 {
			paintingRobot.facing = 1
		} else if x < 0 {
			paintingRobot.facing = 3
		} else if y > 0 {
			paintingRobot.facing = 0
		} else if y < 0 {
			paintingRobot.facing = 2
		}

		paintingRobot.x += x
		paintingRobot.y += y
	}
}
