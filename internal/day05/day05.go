package day05

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type OrderingRule struct {
	Before int
	After  int
}

type Update []int

type PuzzleInput struct {
	Rules   []OrderingRule
	Updates []Update
}

func ParseRule(line string) (OrderingRule, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return OrderingRule{}, fmt.Errorf("invalid rule format: %s", line)
	}

	before, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return OrderingRule{}, fmt.Errorf("invalid before page number: %s", parts[0])
	}

	after, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return OrderingRule{}, fmt.Errorf("invalid after page number: %s", parts[1])
	}

	return OrderingRule{Before: before, After: after}, nil
}

func ParseUpdate(line string) (Update, error) {
	parts := strings.Split(line, ",")
	update := make(Update, len(parts))

	for i, part := range parts {
		page, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid page number: %s", part)
		}
		update[i] = page
	}

	return update, nil
}

func ParseInput(content string) (PuzzleInput, error) {
	sections := strings.Split(content, "\n\n")
	if len(sections) != 2 {
		return PuzzleInput{}, fmt.Errorf("expected 2 sections separated by blank line, got %d", len(sections))
	}

	var rules []OrderingRule
	for _, line := range strings.Split(strings.TrimSpace(sections[0]), "\n") {
		if line == "" {
			continue
		}
		rule, err := ParseRule(line)
		if err != nil {
			return PuzzleInput{}, err
		}
		rules = append(rules, rule)
	}

	var updates []Update
	for _, line := range strings.Split(strings.TrimSpace(sections[1]), "\n") {
		if line == "" {
			continue
		}
		update, err := ParseUpdate(line)
		if err != nil {
			return PuzzleInput{}, err
		}
		updates = append(updates, update)
	}

	return PuzzleInput{Rules: rules, Updates: updates}, nil
}

func IsValidUpdate(update Update, rules []OrderingRule) bool {
	// Create a map for quick lookup of page positions
	pagePos := make(map[int]int)
	for i, page := range update {
		pagePos[page] = i
	}

	// Check each rule against the update
	for _, rule := range rules {
		beforePos, beforeExists := pagePos[rule.Before]
		afterPos, afterExists := pagePos[rule.After]

		// Only check rules where both pages are present in the update
		if beforeExists && afterExists {
			// Rule violation if "before" page comes after "after" page
			if beforePos >= afterPos {
				return false
			}
		}
	}

	return true
}

func GetMiddlePage(update Update) int {
	// For odd-length slices, middle is at index len/2
	return update[len(update)/2]
}

func FixUpdateOrder(update Update, rules []OrderingRule) Update {
	// Create a copy of the update to avoid modifying the original
	fixed := make(Update, len(update))
	copy(fixed, update)

	// Build a map of rules for quick lookup
	ruleMap := make(map[[2]int]bool)
	for _, rule := range rules {
		ruleMap[[2]int{rule.Before, rule.After}] = true
	}

	// Sort using custom comparator based on rules
	sort.Slice(fixed, func(i, j int) bool {
		pageA, pageB := fixed[i], fixed[j]

		// Check if there's a direct rule A|B (A should come before B)
		if ruleMap[[2]int{pageA, pageB}] {
			return true
		}

		// Check if there's a rule B|A (B should come before A)
		if ruleMap[[2]int{pageB, pageA}] {
			return false
		}

		// No direct rule found, maintain original relative order
		return i < j
	})

	return fixed
}

func SolvePart1(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	input, err := ParseInput(string(content))
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, update := range input.Updates {
		if IsValidUpdate(update, input.Rules) {
			sum += GetMiddlePage(update)
		}
	}

	return sum, nil
}

func SolvePart2(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	input, err := ParseInput(string(content))
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, update := range input.Updates {
		if !IsValidUpdate(update, input.Rules) {
			fixed := FixUpdateOrder(update, input.Rules)
			sum += GetMiddlePage(fixed)
		}
	}

	return sum, nil
}
