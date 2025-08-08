package day02

import (
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	// Create a temporary test file
	testData := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	reports, err := parseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("parseInput failed: %v", err)
	}

	if len(reports) != 6 {
		t.Errorf("Expected 6 reports, got %d", len(reports))
	}

	// Test first report: 7 6 4 2 1
	if len(reports[0].Levels) != 5 {
		t.Errorf("First report should have 5 levels, got %d", len(reports[0].Levels))
	}
	expected := []int{7, 6, 4, 2, 1}
	for i, level := range reports[0].Levels {
		if level != expected[i] {
			t.Errorf("First report level %d: expected %d, got %d", i, expected[i], level)
		}
	}

	// Test last report: 1 3 6 7 9
	if len(reports[5].Levels) != 5 {
		t.Errorf("Last report should have 5 levels, got %d", len(reports[5].Levels))
	}
	expectedLast := []int{1, 3, 6, 7, 9}
	for i, level := range reports[5].Levels {
		if level != expectedLast[i] {
			t.Errorf("Last report level %d: expected %d, got %d", i, expectedLast[i], level)
		}
	}
}

func TestIsSafe(t *testing.T) {
	tests := []struct {
		name     string
		report   Report
		expected bool
		reason   string
	}{
		{
			name:     "Safe decreasing by 1-2",
			report:   Report{Levels: []int{7, 6, 4, 2, 1}},
			expected: true,
			reason:   "levels are all decreasing by 1 or 2",
		},
		{
			name:     "Unsafe increase of 5",
			report:   Report{Levels: []int{1, 2, 7, 8, 9}},
			expected: false,
			reason:   "2 7 is an increase of 5",
		},
		{
			name:     "Unsafe decrease of 4",
			report:   Report{Levels: []int{9, 7, 6, 2, 1}},
			expected: false,
			reason:   "6 2 is a decrease of 4",
		},
		{
			name:     "Unsafe mixed directions",
			report:   Report{Levels: []int{1, 3, 2, 4, 5}},
			expected: false,
			reason:   "1 3 is increasing but 3 2 is decreasing",
		},
		{
			name:     "Unsafe no change",
			report:   Report{Levels: []int{8, 6, 4, 4, 1}},
			expected: false,
			reason:   "4 4 is neither an increase or a decrease",
		},
		{
			name:     "Safe increasing by 1-3",
			report:   Report{Levels: []int{1, 3, 6, 7, 9}},
			expected: true,
			reason:   "levels are all increasing by 1, 2, or 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.report.IsSafe()
			if result != tt.expected {
				t.Errorf("IsSafe() = %v, want %v (%s)", result, tt.expected, tt.reason)
			}
		})
	}
}

func TestCountSafeReports(t *testing.T) {
	reports := []Report{
		{Levels: []int{7, 6, 4, 2, 1}}, // Safe
		{Levels: []int{1, 2, 7, 8, 9}}, // Unsafe
		{Levels: []int{9, 7, 6, 2, 1}}, // Unsafe
		{Levels: []int{1, 3, 2, 4, 5}}, // Unsafe
		{Levels: []int{8, 6, 4, 4, 1}}, // Unsafe
		{Levels: []int{1, 3, 6, 7, 9}}, // Safe
	}

	result := CountSafeReports(reports)
	expected := 2

	if result != expected {
		t.Errorf("CountSafeReports() = %d, want %d", result, expected)
	}
}

func TestSolvePart1(t *testing.T) {
	result, err := SolvePart1("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart1 failed: %v", err)
	}

	expected := 2
	if result != expected {
		t.Errorf("SolvePart1() = %d, want %d", result, expected)
	}
}

func TestIsSafeWithDampener(t *testing.T) {
	tests := []struct {
		name     string
		report   Report
		expected bool
		reason   string
	}{
		{
			name:     "Safe without removing any level",
			report:   Report{Levels: []int{7, 6, 4, 2, 1}},
			expected: true,
			reason:   "already safe - decreasing by 1-2",
		},
		{
			name:     "Unsafe regardless of which level is removed",
			report:   Report{Levels: []int{1, 2, 7, 8, 9}},
			expected: false,
			reason:   "2->7 increase of 5, 7->8->9 still has large jumps",
		},
		{
			name:     "Unsafe regardless of which level is removed",
			report:   Report{Levels: []int{9, 7, 6, 2, 1}},
			expected: false,
			reason:   "6->2 decrease of 4, removing any level doesn't fix it",
		},
		{
			name:     "Safe by removing the second level",
			report:   Report{Levels: []int{1, 3, 2, 4, 5}},
			expected: true,
			reason:   "removing 3 gives [1, 2, 4, 5] which is all increasing by 1-2",
		},
		{
			name:     "Safe by removing the third level",
			report:   Report{Levels: []int{8, 6, 4, 4, 1}},
			expected: true,
			reason:   "removing first 4 gives [8, 6, 4, 1] which is all decreasing by 2",
		},
		{
			name:     "Safe without removing any level",
			report:   Report{Levels: []int{1, 3, 6, 7, 9}},
			expected: true,
			reason:   "already safe - increasing by 1-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.report.IsSafeWithDampener()
			if result != tt.expected {
				t.Errorf("IsSafeWithDampener() = %v, want %v (%s)", result, tt.expected, tt.reason)
			}
		})
	}
}

func TestCountSafeReportsWithDampener(t *testing.T) {
	reports := []Report{
		{Levels: []int{7, 6, 4, 2, 1}}, // Safe without removing
		{Levels: []int{1, 2, 7, 8, 9}}, // Unsafe regardless
		{Levels: []int{9, 7, 6, 2, 1}}, // Unsafe regardless
		{Levels: []int{1, 3, 2, 4, 5}}, // Safe by removing second level
		{Levels: []int{8, 6, 4, 4, 1}}, // Safe by removing third level
		{Levels: []int{1, 3, 6, 7, 9}}, // Safe without removing
	}

	result := CountSafeReportsWithDampener(reports)
	expected := 4

	if result != expected {
		t.Errorf("CountSafeReportsWithDampener() = %d, want %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	result, err := SolvePart2("example-input.txt")
	if err != nil {
		t.Fatalf("SolvePart2 failed: %v", err)
	}

	expected := 4
	if result != expected {
		t.Errorf("SolvePart2() = %d, want %d", result, expected)
	}
}
