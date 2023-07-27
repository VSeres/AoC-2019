package main

import (
	"strings"
	"testing"
)

var testMap = []byte(
	`#######...#####
#.....#...#...#
#.....#...#...#
......#...#...#
......#...###.#
......#.....#.#
^########...#.#
......#.#...#.#
......#########
........#...#..
....#########..
....#...#......
....#...#......
....#...#......
....#####......`)

const answer = "R 8 R 8 R 4 R 4 R 8 L 6 L 2 R 4 R 4 R 8 R 8 R 8 L 6 L 2"

func TestGetSteps(t *testing.T) {
	input := make([]int, len(testMap))
	for i, v := range testMap {
		input[i] = int(v)
	}
	output := getSteps(input)
	var buff strings.Builder
	for _, str := range output {
		if buff.Len() > 0 {
			buff.WriteByte(' ')
		}
		buff.WriteString(str)
	}
	if buff.String() != answer {
		t.FailNow()
	}
}
