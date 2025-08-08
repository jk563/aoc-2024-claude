// Package day01 solves the Advent of Code 2024 Day 1 puzzle: "Historian Hysteria".
//
// The puzzle involves two lists of location IDs that need to be analyzed in two different ways:
//
// Part 1: Calculate the total distance by pairing up the smallest number from each list,
// then the second smallest, and so on, calculating the absolute difference for each pair
// and summing all distances.
//
// Part 2: Calculate a similarity score by counting how many times each number from the left
// list appears in the right list, multiplying each left number by its frequency in the right
// list, and summing all results.
//
// Example input:
//
//	3   4
//	4   3
//	2   5
//	1   3
//	3   9
//	3   3
//
// Part 1 expected result: 11 (pairs: (1,3), (2,3), (3,4), (3,5), (4,9), (3,3) with distances 2+1+1+2+5+0)
// Part 2 expected result: 31 (3*3 + 4*1 + 2*0 + 1*0 + 3*3 + 3*3 = 9+4+0+0+9+9)
package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid left number: %s", parts[0])
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid right number: %s", parts[1])
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func calculateTotalDistance(left, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, fmt.Errorf("slice length mismatch: left has %d elements, right has %d elements", len(left), len(right))
	}

	total := 0
	for i := 0; i < len(left); i++ {
		total += calculateDistance(left[i], right[i])
	}

	return total, nil
}

// SolvePart1 solves part 1 of the Day 1 puzzle by calculating the total distance
// between paired location IDs. The solution pairs the smallest number from each list,
// then the second smallest, and so on, calculating the absolute difference for each
// pair and returning the sum of all distances.
func SolvePart1(filename string) (int, error) {
	left, right, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	// Create copies for sorting to avoid modifying original slices
	leftSorted := make([]int, len(left))
	rightSorted := make([]int, len(right))
	copy(leftSorted, left)
	copy(rightSorted, right)

	sort.Ints(leftSorted)
	sort.Ints(rightSorted)

	total, err := calculateTotalDistance(leftSorted, rightSorted)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func countFrequencies(list []int) map[int]int {
	frequencies := make(map[int]int)

	for _, num := range list {
		frequencies[num]++
	}

	return frequencies
}

func calculateSimilarityScore(left, right []int) int {
	frequencies := countFrequencies(right)

	totalScore := 0
	for _, num := range left {
		count := frequencies[num]
		totalScore += num * count
	}

	return totalScore
}

// SolvePart2 solves part 2 of the Day 1 puzzle by calculating a similarity score.
// The solution counts how many times each number from the left list appears in the
// right list, multiplies each left number by its frequency in the right list,
// and returns the sum of all results.
func SolvePart2(filename string) (int, error) {
	left, right, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	return calculateSimilarityScore(left, right), nil
}
