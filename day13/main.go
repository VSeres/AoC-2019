package main

import (
	"aoc-2019/intcode"
	"fmt"
	"strings"
	"time"
)

type tileId int

const (
	empty tileId = iota
	wall
	block
	paddle
	ball
)

type point struct {
	x, y int
}
type tile struct {
	tileType tileId
}

func main() {
	code := intcode.ParseFile("input.txt")
	code.WriteMemory(0, 2)
	blockTiels := 0
	maxX := 0
	maxY := 0
	input := 0
	tiles := make(map[point]tile, 64)
	frame := 1
	var ballPos point
	var paddlePos point
	for !code.Stopped {
		if ballPos.x < paddlePos.x {
			input = -1
		} else if ballPos.x > paddlePos.x {
			input = 1
		} else {
			input = 0
		}
		output := code.Execute([]int{input})
		for i := 0; i <= len(output)-3; i += 3 {
			currentTile := tile{
				tileType: tileId(output[i+2]),
			}
			p := point{output[i], output[i+1]}
			if maxX < p.x {
				maxX = p.x
			}
			if maxY < p.y {
				maxY = p.y
			}
			if currentTile.tileType == block {
				blockTiels++
			}
			if currentTile.tileType == ball {
				ballPos = p
			}
			if currentTile.tileType == paddle {
				paddlePos = p
			}
			tiles[p] = currentTile
		}
		fmt.Print("\033[s")
		var display strings.Builder
		display.WriteString(fmt.Sprintf("frame %d\n", frame))
		frame++
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				pixel := " "
				switch tiles[point{x, y}].tileType {
				case 1:
					pixel = "#"
				case 2:
					pixel = "▓"
				case 3:
					pixel = "═"
				case 4:
					pixel = "■"
				}
				display.WriteString(pixel)
			}
			display.WriteByte(10)
		}
		display.WriteString(fmt.Sprintf("Score %d\n", tiles[point{-1, 0}].tileType))
		fmt.Print(display.String(), "\033[u")
		time.Sleep(16 * time.Millisecond)
	}
	fmt.Println("\033[26B", blockTiels)
}
