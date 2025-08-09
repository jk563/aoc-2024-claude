package day06

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// SolvePart2WithWorkers allows specifying the number of workers
func SolvePart2WithWorkers(filename string, numWorkers int) (int, error) {
	grid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	guard, err := findGuard(grid)
	if err != nil {
		return 0, err
	}

	// Get all positions visited in the original patrol path
	patrolPath := getPatrolPath(grid, guard)
	guardStartPos := guard.Position

	// Collect positions to test (excluding starting position)
	var positions []Position
	for pos := range patrolPath {
		if pos != guardStartPos && grid.Cells[pos.Row][pos.Col] != '#' {
			positions = append(positions, pos)
		}
	}

	// Cap workers at number of positions
	if numWorkers > len(positions) {
		numWorkers = len(positions)
	}

	jobs := make(chan Position, len(positions))
	results := make(chan bool, len(positions))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Each worker needs its own copy of the grid to avoid race conditions
			workerGrid := &Grid{Cells: make([][]rune, len(grid.Cells))}
			for i, row := range grid.Cells {
				workerGrid.Cells[i] = make([]rune, len(row))
				copy(workerGrid.Cells[i], row)
			}

			for pos := range jobs {
				// Temporarily place obstacle
				originalCell := workerGrid.Cells[pos.Row][pos.Col]
				workerGrid.Cells[pos.Row][pos.Col] = '#'

				// Test if this creates a loop
				hasLoop := simulatePatrolWithLoopDetection(workerGrid, guard)
				results <- hasLoop

				// Restore original cell
				workerGrid.Cells[pos.Row][pos.Col] = originalCell
			}
		}()
	}

	// Send jobs
	for _, pos := range positions {
		jobs <- pos
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	loopPositions := 0
	for hasLoop := range results {
		if hasLoop {
			loopPositions++
		}
	}

	return loopPositions, nil
}

func TestWorkerOptimization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping worker optimization test in short mode")
	}

	cpus := runtime.NumCPU()
	fmt.Printf("System has %d logical CPUs\n", cpus)
	
	// Test different multiples of CPU count (8x to 16x to find peak)
	multipliers := []int{8, 9, 10, 11, 12, 13, 14, 15, 16}
	
	type result struct {
		workers int
		time    time.Duration
		answer  int
	}
	
	var results []result
	
	for _, mult := range multipliers {
		workers := cpus * mult
		
		fmt.Printf("\nTesting with %d workers (%dx CPUs)...\n", workers, mult)
		
		// Warm up
		SolvePart2WithWorkers("puzzle-input.txt", workers)
		
		// Measure performance
		start := time.Now()
		answer, err := SolvePart2WithWorkers("puzzle-input.txt", workers)
		elapsed := time.Since(start)
		
		if err != nil {
			t.Errorf("Failed with %d workers: %v", workers, err)
			continue
		}
		
		results = append(results, result{workers, elapsed, answer})
		fmt.Printf("Result: %d, Time: %v\n", answer, elapsed)
	}
	
	// Find the fastest
	if len(results) > 0 {
		fmt.Printf("\n=== PERFORMANCE SUMMARY ===\n")
		fmt.Printf("Workers | Multiplier | Time       | Speedup vs 1x CPU\n")
		fmt.Printf("--------|------------|------------|------------------\n")
		
		baseTime := results[0].time // 1x CPU baseline
		bestTime := time.Duration(1<<63 - 1) // Max duration
		bestWorkers := 0
		bestMult := 0
		
		for i, r := range results {
			mult := multipliers[i]
			
			if r.time < bestTime {
				bestTime = r.time
				bestWorkers = r.workers
				bestMult = mult
			}
			
			speedup := float64(baseTime) / float64(r.time)
			
			fmt.Printf("%7d | %10d | %10v | %.2fx\n", r.workers, mult, r.time, speedup)
		}
		
		fmt.Printf("\nOptimal configuration: %d workers (%dx CPUs) - %.2fx faster than baseline\n", 
			bestWorkers, bestMult, float64(baseTime)/float64(bestTime))
	}
}