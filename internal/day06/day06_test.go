package day06

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	grid, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(grid.Cells) != 10 {
		t.Errorf("Expected 10 rows, got %d", len(grid.Cells))
	}
	if len(grid.Cells[0]) != 10 {
		t.Errorf("Expected 10 columns, got %d", len(grid.Cells[0]))
	}
	if grid.Cells[0][4] != '#' {
		t.Errorf("Expected '#' at [0][4], got %c", grid.Cells[0][4])
	}
	if grid.Cells[6][4] != '^' {
		t.Errorf("Expected '^' at [6][4], got %c", grid.Cells[6][4])
	}
}

func TestFindGuard(t *testing.T) {
	grid, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	guard, err := findGuard(grid)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedPos := Position{Row: 6, Col: 4}
	if guard.Position != expectedPos {
		t.Errorf("Expected guard position %v, got %v", expectedPos, guard.Position)
	}
	if guard.Direction != Up {
		t.Errorf("Expected guard direction Up, got %v", guard.Direction)
	}
}

func TestIsInBounds(t *testing.T) {
	grid, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tests := []struct {
		pos      Position
		expected bool
	}{
		{Position{0, 0}, true},
		{Position{9, 9}, true},
		{Position{-1, 0}, false},
		{Position{0, -1}, false},
		{Position{10, 0}, false},
		{Position{0, 10}, false},
	}

	for _, test := range tests {
		result := isInBounds(grid, test.pos)
		if result != test.expected {
			t.Errorf("isInBounds(%v) = %v, expected %v", test.pos, result, test.expected)
		}
	}
}

func TestIsObstacle(t *testing.T) {
	grid, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tests := []struct {
		pos      Position
		expected bool
	}{
		{Position{0, 4}, true},  // First # in grid
		{Position{1, 9}, true},  // Second # in grid
		{Position{6, 1}, true},  // # at row 6
		{Position{0, 0}, false}, // Empty space
		{Position{6, 4}, false}, // Guard position (^)
	}

	for _, test := range tests {
		result := isObstacle(grid, test.pos)
		if result != test.expected {
			t.Errorf("isObstacle(%v) = %v, expected %v", test.pos, result, test.expected)
		}
	}
}

func TestGetNextPosition(t *testing.T) {
	tests := []struct {
		pos       Position
		direction Direction
		expected  Position
	}{
		{Position{5, 5}, Up, Position{4, 5}},
		{Position{5, 5}, Right, Position{5, 6}},
		{Position{5, 5}, Down, Position{6, 5}},
		{Position{5, 5}, Left, Position{5, 4}},
	}

	for _, test := range tests {
		result := getNextPosition(test.pos, test.direction)
		if result != test.expected {
			t.Errorf("getNextPosition(%v, %v) = %v, expected %v",
				test.pos, test.direction, result, test.expected)
		}
	}
}

func TestTurnRight(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  Direction
	}{
		{Up, Right},
		{Right, Down},
		{Down, Left},
		{Left, Up},
	}

	for _, test := range tests {
		result := turnRight(test.direction)
		if result != test.expected {
			t.Errorf("turnRight(%v) = %v, expected %v", test.direction, result, test.expected)
		}
	}
}

func TestSimulatePatrol(t *testing.T) {
	grid, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	guard, err := findGuard(grid)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	visitedCount := simulatePatrol(grid, guard)

	// From the problem example, the guard should visit 41 distinct positions
	expectedCount := 41
	if visitedCount != expectedCount {
		t.Errorf("Expected %d visited positions, got %d", expectedCount, visitedCount)
	}
}

func TestSolvePart1(t *testing.T) {
	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := 41
	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
