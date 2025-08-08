// Package day03 solves the Advent of Code 2024 Day 3 puzzle: "Mull It Over".
//
// The puzzle involves parsing corrupted computer memory to find valid mul(X,Y) instructions
// and calculating the sum of their multiplication results.
//
// Part 1: Find all valid mul(X,Y) instructions where X and Y are numbers, multiply them,
// and sum all results. Invalid characters and malformed instructions should be ignored.
//
// Example input:
//
//	xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
//
// Valid instructions: mul(2,4), mul(5,5), mul(11,8), mul(8,5)
// Result: 2*4 + 5*5 + 11*8 + 8*5 = 8 + 25 + 88 + 40 = 161
package day03

import (
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const (
	// Instruction type constants
	InstructionTypeMul  = "mul"
	InstructionTypeDo   = "do"
	InstructionTypeDont = "don't"

	// Regex patterns
	mulPattern  = `mul\((\d+),(\d+)\)`
	doPattern   = `do\(\)`
	dontPattern = `don't\(\)`
)

var (
	// Compiled regex patterns (cached for performance)
	mulRe  = regexp.MustCompile(mulPattern)
	doRe   = regexp.MustCompile(doPattern)
	dontRe = regexp.MustCompile(dontPattern)
)

// Instruction represents a single instruction found in the corrupted memory
type Instruction struct {
	Type     string // "mul", "do", or "don't"
	Position int    // Position in the input string
	Value    string // The full matched string
}

// parseInput reads the entire file content as corrupted memory
func parseInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// extractAndMultiply takes a valid mul instruction and returns the product
func extractAndMultiply(instruction string) int {
	matches := mulRe.FindStringSubmatch(instruction)

	if len(matches) != 3 {
		return 0
	}

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])

	return x * y
}

// processCorruptedMemory processes corrupted memory and returns sum of all valid multiplications
// This is now implemented using the same infrastructure as Part 2, but ignores conditional instructions
func processCorruptedMemory(input string) int {
	instructions := findAllInstructions(input)
	total := 0

	for _, instruction := range instructions {
		if instruction.Type == InstructionTypeMul {
			total += extractAndMultiply(instruction.Value)
		}
		// Ignore do() and don't() instructions in Part 1
	}

	return total
}

// SolvePart1 reads input file and returns sum of all valid mul instruction results
func SolvePart1(filename string) (int, error) {
	content, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	return processCorruptedMemory(content), nil
}

// findInstructionsByRegex is a helper function to find instructions using a compiled regex
func findInstructionsByRegex(input string, instructionType string, re *regexp.Regexp) []Instruction {
	matches := re.FindAllStringIndex(input, -1)

	var instructions []Instruction
	for _, match := range matches {
		instructions = append(instructions, Instruction{
			Type:     instructionType,
			Position: match[0],
			Value:    input[match[0]:match[1]],
		})
	}
	return instructions
}

// findAllInstructions finds all valid instructions (mul, do, don't) with their positions
func findAllInstructions(input string) []Instruction {
	var instructions []Instruction

	// Find all instruction types using cached compiled regexes
	instructions = append(instructions, findInstructionsByRegex(input, InstructionTypeMul, mulRe)...)
	instructions = append(instructions, findInstructionsByRegex(input, InstructionTypeDo, doRe)...)
	instructions = append(instructions, findInstructionsByRegex(input, InstructionTypeDont, dontRe)...)

	// Sort by position to process in order
	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i].Position < instructions[j].Position
	})

	return instructions
}

// processWithConditionals processes corrupted memory with conditional instructions
func processWithConditionals(input string) int {
	instructions := findAllInstructions(input)
	enabled := true // mul instructions are enabled at the beginning
	total := 0

	for _, instruction := range instructions {
		switch instruction.Type {
		case InstructionTypeDo:
			enabled = true
		case InstructionTypeDont:
			enabled = false
		case InstructionTypeMul:
			if enabled {
				total += extractAndMultiply(instruction.Value)
			}
		}
	}

	return total
}

// SolvePart2 reads input file and returns sum of enabled mul instruction results
func SolvePart2(filename string) (int, error) {
	content, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	return processWithConditionals(content), nil
}
