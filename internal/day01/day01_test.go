package day01

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	expectedLeft := []int{3, 4, 2, 1, 3, 3}
	expectedRight := []int{4, 3, 5, 3, 9, 3}

	left, right, err := parseInput("example-input.txt")
	if err != nil {
		t.Fatalf("parseInput failed: %v", err)
	}

	if len(left) != len(expectedLeft) {
		t.Errorf("Left list length mismatch. Expected %d, got %d", len(expectedLeft), len(left))
	}

	if len(right) != len(expectedRight) {
		t.Errorf("Right list length mismatch. Expected %d, got %d", len(expectedRight), len(right))
	}

	for i, expected := range expectedLeft {
		if left[i] != expected {
			t.Errorf("Left list[%d] mismatch. Expected %d, got %d", i, expected, left[i])
		}
	}

	for i, expected := range expectedRight {
		if right[i] != expected {
			t.Errorf("Right list[%d] mismatch. Expected %d, got %d", i, expected, right[i])
		}
	}
}

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{3, 7, 4},
		{9, 3, 6},
		{1, 3, 2},
		{2, 3, 1},
		{3, 3, 0},
		{3, 4, 1},
		{3, 5, 2},
		{4, 9, 5},
	}

	for _, test := range tests {
		result := calculateDistance(test.a, test.b)
		if result != test.expected {
			t.Errorf("calculateDistance(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalculateTotalDistance(t *testing.T) {
	left := []int{1, 2, 3, 3, 3, 4}
	right := []int{3, 3, 3, 4, 5, 9}
	expected := 11

	result, err := calculateTotalDistance(left, right)
	if err != nil {
		t.Fatalf("calculateTotalDistance failed: %v", err)
	}
	if result != expected {
		t.Errorf("calculateTotalDistance() = %d, expected %d", result, expected)
	}
}

func TestCalculateTotalDistanceError(t *testing.T) {
	left := []int{1, 2, 3}
	right := []int{3, 3}

	_, err := calculateTotalDistance(left, right)
	if err == nil {
		t.Error("calculateTotalDistance should return error for mismatched slice lengths")
	}
}

func TestSolvePart1(t *testing.T) {
	expected := 11

	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart1 failed: %v", err)
	}

	if result != expected {
		t.Errorf("SolvePart1() = %d, expected %d", result, expected)
	}
}

func TestCountFrequencies(t *testing.T) {
	list := []int{4, 3, 5, 3, 9, 3}

	frequencies := countFrequencies(list)

	expectedFreqs := map[int]int{
		4: 1,
		3: 3,
		5: 1,
		9: 1,
	}

	if len(frequencies) != len(expectedFreqs) {
		t.Errorf("Frequency map length mismatch. Expected %d, got %d", len(expectedFreqs), len(frequencies))
	}

	for num, expectedCount := range expectedFreqs {
		if actualCount, exists := frequencies[num]; !exists {
			t.Errorf("Number %d not found in frequency map", num)
		} else if actualCount != expectedCount {
			t.Errorf("Frequency of %d mismatch. Expected %d, got %d", num, expectedCount, actualCount)
		}
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	expected := 31

	result := calculateSimilarityScore(left, right)
	if result != expected {
		t.Errorf("calculateSimilarityScore() = %d, expected %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := 31

	result, err := SolvePart2("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart2 failed: %v", err)
	}

	if result != expected {
		t.Errorf("SolvePart2() = %d, expected %d", result, expected)
	}
}
