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

var max point

func main() {
	code := intcode.ParseFile("input.txt")
	code.WriteMemory(0, 2)
	blockTiels := 0
	input := 0
	tiles := make(map[point]tile, 64)
	fmt.Print("\033[s")
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
		code.SetInput(input)
		output := code.Execute()
		for i := 0; i <= len(output)-3; i += 3 {
			currentTile := tile{
				tileType: tileId(output[i+2]),
			}
			p := point{output[i], output[i+1]}
			if max.x < p.x {
				max.x = p.x
			}
			if max.y < p.y {
				max.y = p.y
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
		display(tiles)
		time.Sleep(16 * time.Millisecond)
	}
	fmt.Println(blockTiels)
}

func display(tiles map[point]tile) {
	var display strings.Builder
	display.WriteString(fmt.Sprintf("Score %d\n", tiles[point{-1, 0}].tileType))
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
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

	fmt.Print("\033[u", display.String())
}
