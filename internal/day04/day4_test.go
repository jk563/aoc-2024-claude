package day04

import (
	"testing"
)

func TestParseGrid(t *testing.T) {
	grid, err := parseGrid("example-input.txt")
	if err != nil {
		t.Fatalf("parseGrid failed: %v", err)
	}

	// Verify grid dimensions
	if len(grid) != 10 {
		t.Errorf("Expected 10 rows, got %d", len(grid))
	}
	if len(grid[0]) != 10 {
		t.Errorf("Expected 10 columns, got %d", len(grid[0]))
	}

	// Verify first row content
	expectedFirstRow := "MMMSXXMASM"
	for i, char := range expectedFirstRow {
		if grid[0][i] != char {
			t.Errorf("Expected grid[0][%d] = %c, got %c", i, char, grid[0][i])
		}
	}

	// Verify last row content
	expectedLastRow := "MXMXAXMASX"
	lastRowIndex := len(grid) - 1
	for i, char := range expectedLastRow {
		if grid[lastRowIndex][i] != char {
			t.Errorf("Expected grid[%d][%d] = %c, got %c", lastRowIndex, i, char, grid[lastRowIndex][i])
		}
	}
}

func TestParseGridNonExistentFile(t *testing.T) {
	_, err := parseGrid("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestIsValidPosition(t *testing.T) {
	grid := [][]rune{
		{'A', 'B', 'C'},
		{'D', 'E', 'F'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 0, true},    // top-left corner
		{1, 2, true},    // bottom-right corner
		{0, 2, true},    // top-right corner
		{1, 0, true},    // bottom-left corner
		{-1, 0, false},  // above grid
		{0, -1, false},  // left of grid
		{2, 0, false},   // below grid
		{0, 3, false},   // right of grid
		{-1, -1, false}, // outside both dimensions
		{2, 3, false},   // outside both dimensions
	}

	for _, test := range tests {
		result := isValidPosition(grid, test.row, test.col)
		if result != test.expected {
			t.Errorf("isValidPosition(%d, %d) = %v, expected %v", test.row, test.col, result, test.expected)
		}
	}
}

func TestCheckStringInDirection(t *testing.T) {
	// Create a larger grid to accommodate all test cases
	grid := [][]rune{
		{'X', 'M', 'A', 'S', 'X', 'M'},
		{'M', 'A', 'S', 'X', 'M', 'A'},
		{'A', 'S', 'X', 'M', 'A', 'S'},
		{'S', 'X', 'M', 'A', 'S', 'X'},
		{'S', 'A', 'M', 'X', 'M', 'A'},
		{'X', 'M', 'A', 'S', 'X', 'M'},
	}

	tests := []struct {
		row, col, deltaRow, deltaCol int
		target                       string
		expected                     bool
		description                  string
	}{
		{0, 0, 0, 1, "XMAS", true, "horizontal right from (0,0)"},
		{0, 0, 1, 1, "XAXA", true, "diagonal down-right from (0,0): X-A-X-A"},
		{0, 3, 0, -1, "SAMX", true, "horizontal left from (0,3)"},
		{3, 0, -1, 0, "SAMX", true, "vertical up from (3,0)"},
		{0, 0, 0, 1, "ABCD", false, "wrong target word"},
		{0, 0, 0, 1, "XM", true, "partial match XM"},
		{6, 0, 0, 1, "M", false, "out of bounds row"},
		{0, 6, 0, 1, "M", false, "out of bounds col"},
	}

	for _, test := range tests {
		result := checkStringInDirection(grid, test.row, test.col, test.deltaRow, test.deltaCol, test.target)
		if result != test.expected {
			t.Errorf("%s: checkStringInDirection(%d,%d,%d,%d,\"%s\") = %v, expected %v",
				test.description, test.row, test.col, test.deltaRow, test.deltaCol, test.target, result, test.expected)
		}
	}
}

func TestFindXMAS(t *testing.T) {
	grid, err := parseGrid("example-input.txt")
	if err != nil {
		t.Fatalf("Failed to read example-input.txt: %v", err)
	}

	result := findXMAS(grid)
	expected := 18

	if result != expected {
		t.Errorf("findXMAS() = %d, expected %d", result, expected)
	}
}

func TestFindXMASSimple(t *testing.T) {
	// Simple test case with known XMAS patterns
	grid := [][]rune{
		{'X', 'M', 'A', 'S'},
		{'S', 'A', 'M', 'X'},
	}

	result := findXMAS(grid)
	expected := 2 // One XMAS horizontal right, one SAMX horizontal right (which is XMAS backward)

	if result != expected {
		t.Errorf("findXMAS() = %d, expected %d", result, expected)
	}
}

func TestSolvePart1(t *testing.T) {
	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart1 failed: %v", err)
	}

	expected := 18
	if result != expected {
		t.Errorf("SolvePart1() = %d, expected %d", result, expected)
	}
}

func TestCheckXPattern(t *testing.T) {
	tests := []struct {
		name                 string
		grid                 [][]rune
		centerRow, centerCol int
		expected             bool
	}{
		{
			name: "valid X-MAS pattern (MAS + MAS)",
			grid: [][]rune{
				{'M', '.', 'S'},
				{'.', 'A', '.'},
				{'M', '.', 'S'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  true,
		},
		{
			name: "valid X-MAS pattern (MAS + SAM)",
			grid: [][]rune{
				{'M', '.', 'M'},
				{'.', 'A', '.'},
				{'S', '.', 'S'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  true,
		},
		{
			name: "valid X-MAS pattern (SAM + MAS)",
			grid: [][]rune{
				{'S', '.', 'S'},
				{'.', 'A', '.'},
				{'M', '.', 'M'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  true,
		},
		{
			name: "valid X-MAS pattern (SAM + SAM)",
			grid: [][]rune{
				{'S', '.', 'M'},
				{'.', 'A', '.'},
				{'S', '.', 'M'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  true,
		},
		{
			name: "invalid pattern - center not A",
			grid: [][]rune{
				{'M', '.', 'S'},
				{'.', 'X', '.'},
				{'M', '.', 'S'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  false,
		},
		{
			name: "invalid pattern - wrong diagonal characters",
			grid: [][]rune{
				{'M', '.', 'S'},
				{'.', 'A', '.'},
				{'X', '.', 'Y'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  false,
		},
		{
			name: "invalid pattern - only one diagonal correct",
			grid: [][]rune{
				{'M', '.', 'X'},
				{'.', 'A', '.'},
				{'S', '.', 'Y'},
			},
			centerRow: 1,
			centerCol: 1,
			expected:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkXPattern(test.grid, test.centerRow, test.centerCol)
			if result != test.expected {
				t.Errorf("checkXPattern() = %v, expected %v", result, test.expected)
			}
		})
	}
}

func TestFindXMASPattern(t *testing.T) {
	grid, err := parseGrid("example-input.txt")
	if err != nil {
		t.Fatalf("Failed to read example-input.txt: %v", err)
	}

	result := findXMASPattern(grid)
	expected := 9

	if result != expected {
		t.Errorf("findXMASPattern() = %d, expected %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	result, err := SolvePart2("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart2 failed: %v", err)
	}

	expected := 9
	if result != expected {
		t.Errorf("SolvePart2() = %d, expected %d", result, expected)
	}
}
