package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	input := parseInput("input.txt")
	pattern := []int{0, 1, 0, -1}
	output := make([]int16, len(input))

	FFT(input, output, pattern)
}

func FFT(input []int16, output []int16, patter []int) {
	for phase := 0; phase < 100; phase++ {
		for i := range input {
			output[i] = applyPatter(input, patter, i+1)
		}
		copy(input, output)
	}
}

func applyPatter(arr []int16, pattern []int, num int) int16 {
	sum := 0
	repeat := num
	mult := pattern[0]
	var patterPoint int
	repeat -= 1
	for _, v := range arr {
		if repeat <= 0 {
			patterPoint++
			index := patterPoint % 4
			mult = pattern[index]
			repeat = num
		}
		// fmt.Printf("%d*%d + ", mult, v)
		sum += mult * int(v)
		repeat--
	}

	if sum < 0 {
		sum = -sum
	}
	// fmt.Println(" = ", sum%10)
	return int16(sum % 10)
}

func parseInput(name string) []int16 {
	file, _ := os.Open(name)
	reader := bufio.NewReader(file)
	input := make([]int16, 0, 1024)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		num, _ := strconv.Atoi(string(char))
		input = append(input, int16(num))
	}
	return input
}
