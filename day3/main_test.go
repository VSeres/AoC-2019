package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	wires := parseFile("./example2.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(wires)
	}
}

func BenchmarkSolveOld(b *testing.B) {
	wires := parseFile("./example2.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOld(wires)
	}
}

func TestSolve(t *testing.T) {
	wiresOne := parseFile("example1.txt")
	wiresTwo := parseFile("example2.txt")

	parOne, partTwo := solve(wiresOne)
	if parOne != 159 || partTwo != 610 {
		t.Errorf("(159, 610) should be (%d, %d)", parOne, partTwo)
	}

	parOne, partTwo = solve(wiresTwo)
	if parOne != 135 || partTwo != 410 {
		t.Errorf("(135, 410) should be (%d, %d)", parOne, partTwo)
	}
}
