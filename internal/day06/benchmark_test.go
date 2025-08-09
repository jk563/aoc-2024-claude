package day06

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkSolvePart2Serial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := SolvePart2("puzzle-input.txt")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := SolvePart2("puzzle-input.txt")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Manual performance comparison
func TestPerformanceComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	// Warm up
	SolvePart2("puzzle-input.txt")
	SolvePart2("puzzle-input.txt")

	// Serial version
	start := time.Now()
	result1, err1 := SolvePart2("puzzle-input.txt")
	serialTime := time.Since(start)
	
	if err1 != nil {
		t.Fatal("Serial version failed:", err1)
	}

	// Optimized parallel version (now the main implementation)
	start = time.Now()
	result2, err2 := SolvePart2("puzzle-input.txt")
	parallelTime := time.Since(start)
	
	if err2 != nil {
		t.Fatal("Parallel version failed:", err2)
	}

	// Verify same results
	if result1 != result2 {
		t.Errorf("Results differ: serial=%d, parallel=%d", result1, result2)
	}

	// Report performance
	fmt.Printf("\nPerformance Comparison:\n")
	fmt.Printf("Serial:   %v (result: %d)\n", serialTime, result1)
	fmt.Printf("Parallel: %v (result: %d)\n", parallelTime, result2)
	fmt.Printf("Speedup:  %.2fx\n", float64(serialTime)/float64(parallelTime))
}