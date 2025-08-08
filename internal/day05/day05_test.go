package day05

import (
	"os"
	"testing"
)

func TestParseRule(t *testing.T) {
	tests := []struct {
		input    string
		expected OrderingRule
	}{
		{"47|53", OrderingRule{Before: 47, After: 53}},
		{"97|13", OrderingRule{Before: 97, After: 13}},
		{"75|29", OrderingRule{Before: 75, After: 29}},
	}

	for _, test := range tests {
		result, err := ParseRule(test.input)
		if err != nil {
			t.Errorf("ParseRule(%q) returned error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("ParseRule(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestParseUpdate(t *testing.T) {
	tests := []struct {
		input    string
		expected Update
	}{
		{"75,47,61,53,29", Update{75, 47, 61, 53, 29}},
		{"97,61,53,29,13", Update{97, 61, 53, 29, 13}},
		{"75,29,13", Update{75, 29, 13}},
	}

	for _, test := range tests {
		result, err := ParseUpdate(test.input)
		if err != nil {
			t.Errorf("ParseUpdate(%q) returned error: %v", test.input, err)
		}
		if len(result) != len(test.expected) {
			t.Errorf("ParseUpdate(%q) length = %d, expected %d", test.input, len(result), len(test.expected))
			continue
		}
		for i, page := range test.expected {
			if result[i] != page {
				t.Errorf("ParseUpdate(%q)[%d] = %d, expected %d", test.input, i, result[i], page)
			}
		}
	}
}

func TestParseInputWithExampleFile(t *testing.T) {
	content, err := os.ReadFile("example-input.txt")
	if err != nil {
		t.Skip("example-input.txt not found, skipping test")
	}

	result, err := ParseInput(string(content))
	if err != nil {
		t.Errorf("ParseInput returned error: %v", err)
	}

	// Based on the example in the problem description
	expectedRuleCount := 21
	expectedUpdateCount := 6

	if len(result.Rules) != expectedRuleCount {
		t.Errorf("Expected %d rules, got %d", expectedRuleCount, len(result.Rules))
	}
	if len(result.Updates) != expectedUpdateCount {
		t.Errorf("Expected %d updates, got %d", expectedUpdateCount, len(result.Updates))
	}

	// Verify some specific rules and updates from the example
	expectedFirstRule := OrderingRule{Before: 47, After: 53}
	if result.Rules[0] != expectedFirstRule {
		t.Errorf("First rule = %v, expected %v", result.Rules[0], expectedFirstRule)
	}

	expectedFirstUpdate := Update{75, 47, 61, 53, 29}
	if len(result.Updates[0]) != len(expectedFirstUpdate) {
		t.Errorf("First update length = %d, expected %d", len(result.Updates[0]), len(expectedFirstUpdate))
	} else {
		for i, page := range expectedFirstUpdate {
			if result.Updates[0][i] != page {
				t.Errorf("First update[%d] = %d, expected %d", i, result.Updates[0][i], page)
			}
		}
	}
}

func TestIsValidUpdate(t *testing.T) {
	rules := []OrderingRule{
		{Before: 47, After: 53},
		{Before: 97, After: 75},
		{Before: 29, After: 13},
	}

	tests := []struct {
		update   Update
		expected bool
		desc     string
	}{
		{Update{47, 53}, true, "valid - 47 before 53"},
		{Update{53, 47}, false, "invalid - 53 before 47 violates 47|53"},
		{Update{97, 75}, true, "valid - 97 before 75"},
		{Update{75, 97}, false, "invalid - 75 before 97 violates 97|75"},
		{Update{29, 13}, true, "valid - 29 before 13"},
		{Update{13, 29}, false, "invalid - 13 before 29 violates 29|13"},
		{Update{47, 53, 29}, true, "valid - no applicable rules violated"},
		{Update{1, 2, 3}, true, "valid - no applicable rules"},
	}

	for _, test := range tests {
		result := IsValidUpdate(test.update, rules)
		if result != test.expected {
			t.Errorf("IsValidUpdate(%v) = %v, expected %v (%s)", test.update, result, test.expected, test.desc)
		}
	}
}

func TestGetMiddlePage(t *testing.T) {
	tests := []struct {
		update   Update
		expected int
	}{
		{Update{75, 47, 61, 53, 29}, 61},
		{Update{97, 61, 53, 29, 13}, 53},
		{Update{75, 29, 13}, 29},
		{Update{1}, 1},
		{Update{1, 2, 3}, 2},
	}

	for _, test := range tests {
		result := GetMiddlePage(test.update)
		if result != test.expected {
			t.Errorf("GetMiddlePage(%v) = %d, expected %d", test.update, result, test.expected)
		}
	}
}

func TestSolvePart1WithExampleFile(t *testing.T) {
	content, err := os.ReadFile("example-input.txt")
	if err != nil {
		t.Skip("example-input.txt not found, skipping test")
	}

	input, err := ParseInput(string(content))
	if err != nil {
		t.Errorf("ParseInput returned error: %v", err)
		return
	}

	// Test the logic directly with parsed input
	sum := 0
	for _, update := range input.Updates {
		if IsValidUpdate(update, input.Rules) {
			sum += GetMiddlePage(update)
		}
	}

	expected := 143 // From problem description: 61 + 53 + 29 = 143
	if sum != expected {
		t.Errorf("Part1 logic = %d, expected %d", sum, expected)
	}
}

func TestFixUpdateOrder(t *testing.T) {
	rules := []OrderingRule{
		{Before: 47, After: 53},
		{Before: 97, After: 13},
		{Before: 97, After: 61},
		{Before: 97, After: 47},
		{Before: 75, After: 29},
		{Before: 61, After: 13},
		{Before: 75, After: 53},
		{Before: 29, After: 13},
		{Before: 97, After: 29},
		{Before: 53, After: 29},
		{Before: 61, After: 53},
		{Before: 97, After: 53},
		{Before: 61, After: 29},
		{Before: 47, After: 13},
		{Before: 75, After: 47},
		{Before: 97, After: 75},
		{Before: 47, After: 61},
		{Before: 75, After: 61},
		{Before: 47, After: 29},
		{Before: 75, After: 13},
		{Before: 53, After: 13},
	}

	tests := []struct {
		input    Update
		expected Update
		desc     string
	}{
		{Update{75, 97, 47, 61, 53}, Update{97, 75, 47, 61, 53}, "first example case"},
		{Update{61, 13, 29}, Update{61, 29, 13}, "second example case"},
		{Update{97, 13, 75, 29, 47}, Update{97, 75, 47, 29, 13}, "third example case"},
	}

	for _, test := range tests {
		result := FixUpdateOrder(test.input, rules)
		if len(result) != len(test.expected) {
			t.Errorf("FixUpdateOrder(%v) length = %d, expected %d (%s)", test.input, len(result), len(test.expected), test.desc)
			continue
		}
		for i, page := range test.expected {
			if result[i] != page {
				t.Errorf("FixUpdateOrder(%v)[%d] = %d, expected %d (%s)", test.input, i, result[i], page, test.desc)
			}
		}
	}
}

func TestSolvePart2WithExampleFile(t *testing.T) {
	content, err := os.ReadFile("example-input.txt")
	if err != nil {
		t.Skip("example-input.txt not found, skipping test")
	}

	input, err := ParseInput(string(content))
	if err != nil {
		t.Errorf("ParseInput returned error: %v", err)
		return
	}

	// Test the logic directly with parsed input
	sum := 0
	for _, update := range input.Updates {
		if !IsValidUpdate(update, input.Rules) {
			fixed := FixUpdateOrder(update, input.Rules)
			sum += GetMiddlePage(fixed)
		}
	}

	expected := 123 // From problem description: 47 + 29 + 47 = 123
	if sum != expected {
		t.Errorf("Part2 logic = %d, expected %d", sum, expected)
	}
}
