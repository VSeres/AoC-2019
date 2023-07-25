package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var offset int

func main() {
	input := parseInput("input.txt")
	partOne := make([]int16, len(input))
	copy(partOne, input)
	fmt.Println("Part one: ", FFT(partOne))

	for i := 0; i < 7; i++ {
		offset += int(math.Pow10(6-i)) * int(input[i])
	}
	partTwo := make([]int16, len(input)*10000)
	for i := 0; i < len(partTwo); i++ {
		partTwo[i] = input[i%len(input)]
	}
	fmt.Println("part two: ", FFT2(partTwo[offset:]))
}

func FFT(input []int16) string {
	output := make([]int16, len(input))
	for phase := 0; phase < 100; phase++ {
		for i := range input {
			output[i] = applyPatter(input, i+1)
		}
		copy(input, output)
	}
	var buff strings.Builder
	for i := 0; i < 8; i++ {
		buff.WriteByte(byte(output[i]) + 48)
	}
	return buff.String()
}

var pattern = []int{0, 1, 0, -1}

func applyPatter(arr []int16, num int) int16 {
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
		sum += mult * int(v)
		repeat--
	}

	if sum < 0 {
		sum = -sum
	}
	return int16(sum % 10)
}

func FFT2(arr []int16) string {
	out := make([]int16, len(arr))
	for phase := 0; phase < 100; phase++ {
		sum := 0
		for j := 0; j < len(arr); j++ {
			sum += int(arr[j])
		}
		for i := 0; i < len(arr); i++ {
			out[i] = int16((sum % 10))
			sum -= int(arr[i])
		}
		copy(arr, out)
	}
	var buff strings.Builder
	for i := 0; i < 8; i++ {
		buff.WriteByte(byte(out[i]) + 48)
	}
	return buff.String()
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
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println(err, char)
		}
		input = append(input, int16(num))
	}
	return input
}
