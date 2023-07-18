package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const (
	imageHeight = 6
	imageWidth  = 25
)

type minimum struct {
	zeroCount int
	oneCount  int
	twoCount  int
}

func main() {
	layers := parseFile("input.txt")
	minZero := minimum{
		zeroCount: math.MaxInt,
	}
	image := make([][]byte, imageHeight)
	for i := range image {
		image[i] = make([]byte, imageWidth)
	}

	for _, layer := range layers {
		count := minimum{}
		for pos, digit := range layer {
			if digit == '0' {
				count.zeroCount += 1
			} else if digit == '1' {
				count.oneCount += 1
			} else if digit == '2' {
				count.twoCount += 1
			}
			y := pos / (imageWidth)
			x := pos % imageWidth
			pixel := image[y][x]
			if pixel == 0 || pixel == '2' {
				image[y][x] = digit
			}
		}
		if minZero.zeroCount > count.zeroCount {
			minZero = count
		}
	}
	fmt.Println(minZero.twoCount * minZero.oneCount)
	for _, row := range image {
		for _, v := range row {
			if v == '1' {
				fmt.Printf("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func parseFile(name string) [][]byte {
	file, _ := os.Open(name)
	reader := bufio.NewReader(file)
	layers := make([][]byte, 0)
	var err error
	var char rune
	for err != io.EOF {
		layer := make([]byte, imageHeight*imageWidth)
		for i := range layer {
			char, _, _ = reader.ReadRune()
			layer[i] = byte(char)
		}
		layers = append(layers, layer)
		_, err = reader.Peek(1)
	}
	return layers
}
