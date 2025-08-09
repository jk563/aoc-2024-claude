package day07

import (
	"testing"
)


// Phase 3.1: Test evaluating expression left-to-right
func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		operands  []int
		operators []string
		expected  int
	}{
		{
			operands:  []int{10, 19},
			operators: []string{"*"},
			expected:  190,
		},
		{
			operands:  []int{81, 40, 27},
			operators: []string{"+", "*"},
			expected:  3267, // (81 + 40) * 27 = 121 * 27 = 3267
		},
		{
			operands:  []int{81, 40, 27},
			operators: []string{"*", "+"},
			expected:  3267, // (81 * 40) + 27 = 3240 + 27 = 3267
		},
		{
			operands:  []int{11, 6, 16, 20},
			operators: []string{"+", "*", "+"},
			expected:  292, // ((11 + 6) * 16) + 20 = (17 * 16) + 20 = 272 + 20 = 292
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := evaluateExpression(tt.operands, tt.operators)
			if result != tt.expected {
				t.Errorf("evaluateExpression(%v, %v) = %d, want %d", tt.operands, tt.operators, result, tt.expected)
			}
		})
	}
}

// Test checking if equation can be made valid with Part 1 operators
func TestCanSolveEquationPart1(t *testing.T) {
	tests := []struct {
		testValue int
		operands  []int
		expected  bool
	}{
		{190, []int{10, 19}, true},        // 10 * 19 = 190
		{3267, []int{81, 40, 27}, true},   // Multiple solutions work
		{83, []int{17, 5}, false},         // No valid combination
		{156, []int{15, 6}, false},        // No valid combination  
		{292, []int{11, 6, 16, 20}, true}, // 11 + 6 * 16 + 20 = 292
	}

	part1Operators := []string{"+", "*"}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := canSolveEquation(tt.testValue, tt.operands, part1Operators)
			if result != tt.expected {
				t.Errorf("canSolveEquation(%d, %v, %v) = %v, want %v", tt.testValue, tt.operands, part1Operators, result, tt.expected)
			}
		})
	}
}

// Phase 5.1: Test calculating total calibration result using example file
func TestSolvePart1(t *testing.T) {
	expected := 3749 // 190 + 3267 + 292 (only solvable equations)
	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart1 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("SolvePart1() = %d, want %d", result, expected)
	}
}

// Test concatenation operation
func TestConcatenateNumbers(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 345, 12345},
		{15, 6, 156},
		{17, 8, 178},
		{6, 8, 68},
		{48, 6, 486},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := concatenateNumbersMath(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("concatenateNumbersMath(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Test evaluating expressions with concatenation operator
func TestEvaluateExpressionWithConcatenation(t *testing.T) {
	tests := []struct {
		operands  []int
		operators []string
		expected  int
	}{
		{
			operands:  []int{15, 6},
			operators: []string{"||"},
			expected:  156,
		},
		{
			operands:  []int{17, 8, 14},
			operators: []string{"||", "+"},
			expected:  192, // 17 || 8 + 14 = 178 + 14 = 192
		},
		{
			operands:  []int{6, 8, 6, 15},
			operators: []string{"*", "||", "*"},
			expected:  7290, // 6 * 8 || 6 * 15 = 48 || 6 * 15 = 486 * 15 = 7290
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := evaluateExpression(tt.operands, tt.operators)
			if result != tt.expected {
				t.Errorf("evaluateExpression(%v, %v) = %d, want %d", tt.operands, tt.operators, result, tt.expected)
			}
		})
	}
}

// Test Part 2 equation solving
func TestCanSolveEquationPart2(t *testing.T) {
	tests := []struct {
		testValue int
		operands  []int
		expected  bool
	}{
		{156, []int{15, 6}, true},         // 15 || 6 = 156
		{7290, []int{6, 8, 6, 15}, true}, // 6 * 8 || 6 * 15 = 7290
		{192, []int{17, 8, 14}, true},    // 17 || 8 + 14 = 192
	}

	part2Operators := []string{"+", "*", "||"}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := canSolveEquation(tt.testValue, tt.operands, part2Operators)
			if result != tt.expected {
				t.Errorf("canSolveEquation(%d, %v, %v) = %v, want %v", tt.testValue, tt.operands, part2Operators, result, tt.expected)
			}
		})
	}
}

// Test Part 2 solution
func TestSolvePart2(t *testing.T) {
	expected := 11387 // All 6 solvable equations
	result, err := SolvePart2("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart2 returned error: %v", err)
	}
	if result != expected {
		t.Errorf("SolvePart2() = %d, want %d", result, expected)
	}
}

// Test mathematical concatenation produces correct results
func TestMathematicalConcatenation(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 345, 12345},
		{15, 6, 156},
		{17, 8, 178},
		{6, 8, 68},
		{48, 6, 486},
		{123, 4567, 1234567},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			mathResult := concatenateNumbersMath(tt.a, tt.b)
			
			if mathResult != tt.expected {
				t.Errorf("concatenateNumbersMath(%d, %d) = %d, want %d", tt.a, tt.b, mathResult, tt.expected)
			}
		})
	}
}


