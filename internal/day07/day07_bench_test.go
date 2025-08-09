package day07

import (
	"fmt"
	"runtime"
	"testing"
)

// Benchmark Part 1 solution
func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := SolvePart1("puzzle-input.txt")
		if err != nil {
			b.Fatalf("SolvePart1 failed: %v", err)
		}
	}
}

// Benchmark Part 2 solution
func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := SolvePart2("puzzle-input.txt")
		if err != nil {
			b.Fatalf("SolvePart2 failed: %v", err)
		}
	}
}

// Benchmark concatenation function
func BenchmarkConcatenateNumbers_Math(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concatenateNumbersMath(12345, 6789)
	}
}


// Benchmark Part 1 with different worker counts (1x and 3x NumCPU)
func BenchmarkSolvePart1WorkerCounts(b *testing.B) {
	equations, err := parseInput("puzzle-input.txt")
	if err != nil {
		b.Fatalf("Failed to parse input: %v", err)
	}
	part1Operators := []string{"+", "*"}
	numCPU := runtime.NumCPU()
	
	for _, multiple := range []int{1, 3} {
		numWorkers := numCPU * multiple
		b.Run(fmt.Sprintf("%dx_CPU_%d_workers", multiple, numWorkers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = solveEquationsParallelWithWorkers(equations, part1Operators, numWorkers)
			}
		})
	}
}

// Benchmark Part 2 with different worker counts (1x and 3x NumCPU)
func BenchmarkSolvePart2WorkerCounts(b *testing.B) {
	equations, err := parseInput("puzzle-input.txt")
	if err != nil {
		b.Fatalf("Failed to parse input: %v", err)
	}
	part2Operators := []string{"+", "*", "||"}
	numCPU := runtime.NumCPU()
	
	for _, multiple := range []int{1, 3} {
		numWorkers := numCPU * multiple
		b.Run(fmt.Sprintf("%dx_CPU_%d_workers", multiple, numWorkers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = solveEquationsParallelWithWorkers(equations, part2Operators, numWorkers)
			}
		})
	}
}

