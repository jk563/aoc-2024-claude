// Package day04 solves the Advent of Code 2024 Day 4 puzzle: "Ceres Search".
//
// The puzzle involves finding all occurrences of the word "XMAS" in a word search grid.
// XMAS can appear in any of 8 directions: horizontal (left/right), vertical (up/down),
// and diagonal (4 directions). Words can be written backwards and can overlap.
//
// Part 1: Count all occurrences of "XMAS" in the grid.
// Part 2: [To be implemented]
//
// Example input:
//
//	MMMSXXMASM
//	MSAMXMSMSA
//	AMXSXMAAMM
//	...
//
// Part 1 expected result: 18
package day04

import (
	"bufio"
	"os"
)

func parseGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func isValidPosition(grid [][]rune, row, col int) bool {
	if len(grid) == 0 {
		return false
	}
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func checkStringInDirection(grid [][]rune, row, col, deltaRow, deltaCol int, target string) bool {
	for i, char := range target {
		newRow := row + i*deltaRow
		newCol := col + i*deltaCol

		if !isValidPosition(grid, newRow, newCol) {
			return false
		}

		if grid[newRow][newCol] != char {
			return false
		}
	}
	return true
}

func findXMAS(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	target := "XMAS"

	// All 8 directions: right, down, diagonal down-right, diagonal down-left,
	// left, up, diagonal up-left, diagonal up-right
	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{1, -1},  // diagonal down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // diagonal up-left
		{-1, 1},  // diagonal up-right
	}

	// Check each position in the grid
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			// Check each direction from this position
			for _, dir := range directions {
				if checkStringInDirection(grid, row, col, dir[0], dir[1], target) {
					count++
				}
			}
		}
	}

	return count
}

// SolvePart1 solves part 1 of the Day 4 puzzle by finding all occurrences of "XMAS"
// in the word search grid. The solution searches in all 8 directions from each position
// and counts all matches.
func SolvePart1(filename string) (int, error) {
	grid, err := parseGrid(filename)
	if err != nil {
		return 0, err
	}

	return findXMAS(grid), nil
}

func checkXPattern(grid [][]rune, centerRow, centerCol int) bool {
	// Check if center position is valid and contains 'A'
	if !isValidPosition(grid, centerRow, centerCol) || grid[centerRow][centerCol] != 'A' {
		return false
	}

	// Check both diagonals using checkStringInDirection
	// Diagonal 1: top-left to bottom-right (should be "MAS" or "SAM")
	topLeftRow, topLeftCol := centerRow-1, centerCol-1
	diagonal1Valid := checkStringInDirection(grid, topLeftRow, topLeftCol, 1, 1, "MAS") ||
		checkStringInDirection(grid, topLeftRow, topLeftCol, 1, 1, "SAM")

	// Diagonal 2: top-right to bottom-left (should be "MAS" or "SAM")
	topRightRow, topRightCol := centerRow-1, centerCol+1
	diagonal2Valid := checkStringInDirection(grid, topRightRow, topRightCol, 1, -1, "MAS") ||
		checkStringInDirection(grid, topRightRow, topRightCol, 1, -1, "SAM")

	return diagonal1Valid && diagonal2Valid
}

func findXMASPattern(grid [][]rune) int {
	if len(grid) < 3 {
		return 0
	}

	count := 0

	// Check each possible center position (must have room for the X pattern)
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			if checkXPattern(grid, row, col) {
				count++
			}
		}
	}

	return count
}

// SolvePart2 solves part 2 of the Day 4 puzzle by finding X-MAS patterns.
// An X-MAS pattern consists of two "MAS" words arranged in an X shape,
// where the 'A' is at the center and each "MAS" can be written forwards or backwards.
func SolvePart2(filename string) (int, error) {
	grid, err := parseGrid(filename)
	if err != nil {
		return 0, err
	}

	return findXMASPattern(grid), nil
}
