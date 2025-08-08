package day03

import (
	"testing"
)

func TestExtractAndMultiply(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple multiplication",
			input:    "mul(2,4)",
			expected: 8,
		},
		{
			name:     "single digits",
			input:    "mul(5,5)",
			expected: 25,
		},
		{
			name:     "larger numbers",
			input:    "mul(11,8)",
			expected: 88,
		},
		{
			name:     "very large numbers",
			input:    "mul(12345,67890)",
			expected: 838102050,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractAndMultiply(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestProcessCorruptedMemory(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example from problem",
			input:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expected: 161, // 2*4 + 5*5 + 11*8 + 8*5 = 8 + 25 + 88 + 40 = 161
		},
		{
			name:     "single mul",
			input:    "mul(10,20)",
			expected: 200,
		},
		{
			name:     "no valid muls",
			input:    "hello world mul(invalid",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processCorruptedMemory(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 161
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestFindAllInstructions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Instruction
	}{
		{
			name:  "mul only",
			input: "mul(2,4)",
			expected: []Instruction{
				{Type: "mul", Position: 0, Value: "mul(2,4)"},
			},
		},
		{
			name:  "do and don't only",
			input: "do()don't()",
			expected: []Instruction{
				{Type: "do", Position: 0, Value: "do()"},
				{Type: "don't", Position: 4, Value: "don't()"},
			},
		},
		{
			name:  "part 2 example",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: []Instruction{
				{Type: "mul", Position: 1, Value: "mul(2,4)"},
				{Type: "don't", Position: 20, Value: "don't()"},
				{Type: "mul", Position: 28, Value: "mul(5,5)"},
				{Type: "mul", Position: 48, Value: "mul(11,8)"},
				{Type: "do", Position: 59, Value: "do()"},
				{Type: "mul", Position: 64, Value: "mul(8,5)"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAllInstructions(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d instructions, got %d", len(tt.expected), len(result))
				return
			}
			for i, expected := range tt.expected {
				if result[i].Type != expected.Type {
					t.Errorf("instruction %d: expected type %s, got %s", i, expected.Type, result[i].Type)
				}
				if result[i].Position != expected.Position {
					t.Errorf("instruction %d: expected position %d, got %d", i, expected.Position, result[i].Position)
				}
				if result[i].Value != expected.Value {
					t.Errorf("instruction %d: expected value %s, got %s", i, expected.Value, result[i].Value)
				}
			}
		})
	}
}

func TestProcessWithConditionals(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "part 2 example",
			input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expected: 48, // 2*4 + 8*5 = 8 + 40 = 48
		},
		{
			name:     "all enabled",
			input:    "mul(2,3)mul(4,5)",
			expected: 26, // 2*3 + 4*5 = 6 + 20 = 26
		},
		{
			name:     "all disabled",
			input:    "don't()mul(2,3)mul(4,5)",
			expected: 0,
		},
		{
			name:     "enable then disable",
			input:    "mul(2,3)don't()mul(4,5)do()mul(6,7)",
			expected: 48, // 2*3 + 6*7 = 6 + 42 = 48
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processWithConditionals(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	result, err := SolvePart2("example-part2-input.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 48
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
