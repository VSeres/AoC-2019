package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	wires := parseFile("./example.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(wires)
	}
}

func BenchmarkSolveOld(b *testing.B) {
	wires := parseFile("./example.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveOld(wires)
	}
}
