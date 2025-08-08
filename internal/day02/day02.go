// Package day02 solves the Advent of Code 2024 Day 2 puzzle: "Red-Nosed Reports".
//
// The puzzle involves analyzing reports that contain sequences of levels (numbers) to determine
// which reports are "safe" according to specific safety criteria:
//
// Part 1: A report is safe if both conditions are met:
//  1. All levels must be either all increasing or all decreasing (monotonic)
//  2. Adjacent levels must differ by at least 1 and at most 3
//
// Part 2: A report is safe if it meets the Part 1 criteria OR becomes safe after removing
// exactly one level (Problem Dampener). This allows tolerance for a single problematic level.
//
// Example input:
//
//	7 6 4 2 1
//	1 2 7 8 9
//	9 7 6 2 1
//	1 3 2 4 5
//	8 6 4 4 1
//	1 3 6 7 9
//
// Part 1 analysis:
//   - 7 6 4 2 1: Safe (decreasing by 1-2 each step)
//   - 1 2 7 8 9: Unsafe (increase of 5 between 2 and 7)
//   - 9 7 6 2 1: Unsafe (decrease of 4 between 6 and 2)
//   - 1 3 2 4 5: Unsafe (mixed directions: up, down, up, up)
//   - 8 6 4 4 1: Unsafe (no change between the two 4s)
//   - 1 3 6 7 9: Safe (increasing by 1-3 each step)
//
// Part 1 expected result: 2 safe reports
// Part 2 expected result: 4 safe reports (with Problem Dampener tolerance)
package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Report represents a single report containing levels
type Report struct {
	Levels []int
}

// parseInput reads a file and converts each line to a Report
func parseInput(filename string) ([]Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports []Report
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		levels := make([]int, 0, len(fields))

		for _, field := range fields {
			level, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			levels = append(levels, level)
		}

		reports = append(reports, Report{Levels: levels})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

// IsSafe checks if a report is safe according to the rules:
// 1. All levels must be either increasing or decreasing
// 2. Adjacent levels must differ by at least 1 and at most 3
func (r Report) IsSafe() bool {
	if len(r.Levels) < 2 {
		return true
	}

	return r.isMonotonic() && r.hasValidDifferences()
}

// isMonotonic checks if the sequence is either all increasing or all decreasing
func (r Report) isMonotonic() bool {
	if len(r.Levels) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(r.Levels); i++ {
		if r.Levels[i] > r.Levels[i-1] {
			decreasing = false
		} else if r.Levels[i] < r.Levels[i-1] {
			increasing = false
		} else {
			// Equal values mean neither increasing nor decreasing
			return false
		}
	}

	return increasing || decreasing
}

// hasValidDifferences checks if adjacent levels differ by 1-3
func (r Report) hasValidDifferences() bool {
	for i := 1; i < len(r.Levels); i++ {
		diff := r.Levels[i] - r.Levels[i-1]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

// CountSafeReports counts how many reports in the slice are safe
func CountSafeReports(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafe() {
			count++
		}
	}
	return count
}

// IsSafeWithDampener checks if a report is safe with Problem Dampener
// (safe as-is OR safe after removing exactly one level)
func (r Report) IsSafeWithDampener() bool {
	// First check if already safe
	if r.IsSafe() {
		return true
	}

	// Try removing each level one at a time
	for i := 0; i < len(r.Levels); i++ {
		// Create new slice without element at index i
		dampened := make([]int, 0, len(r.Levels)-1)
		dampened = append(dampened, r.Levels[:i]...)
		dampened = append(dampened, r.Levels[i+1:]...)

		// Check if the dampened report is safe
		dampenedReport := Report{Levels: dampened}
		if dampenedReport.IsSafe() {
			return true
		}
	}

	return false
}

// CountSafeReportsWithDampener counts how many reports are safe with Problem Dampener
func CountSafeReportsWithDampener(reports []Report) int {
	count := 0
	for _, report := range reports {
		if report.IsSafeWithDampener() {
			count++
		}
	}
	return count
}

// SolvePart1 reads input file and returns the count of safe reports
func SolvePart1(filename string) (int, error) {
	reports, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	return CountSafeReports(reports), nil
}

// SolvePart2 reads input file and returns the count of safe reports with Problem Dampener
func SolvePart2(filename string) (int, error) {
	reports, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	return CountSafeReportsWithDampener(reports), nil
}
