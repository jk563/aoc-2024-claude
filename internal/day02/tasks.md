# Day 2: Red-Nosed Reports - Implementation Plan

## Problem Summary
- Parse reports containing space-separated numbers (levels)
- Determine if each report is "safe" based on two rules:
  1. All levels must be either increasing or decreasing (no mixed directions)
  2. Adjacent levels must differ by at least 1 and at most 3

## Implementation Tasks

### Phase 1: Core Data Structure and Parsing
- [x] Create Report struct to represent a single report
- [x] Implement ParseInput function to read file and convert to []Report
- [x] Write unit tests for ParseInput with example data

### Phase 2: Safety Validation Logic
- [x] Implement IsSafe method on Report struct
- [x] Create helper function to check if sequence is monotonic (all increasing or all decreasing)
- [x] Create helper function to validate adjacent differences (1-3 range)
- [x] Write comprehensive unit tests for IsSafe with all example cases

### Phase 3: Main Solution Function
- [x] Implement CountSafeReports function that counts safe reports in a slice
- [x] Write unit tests for CountSafeReports using example data
- [x] Ensure all tests pass

### Phase 4: Integration and CLI
- [x] Create SolvePart1 function that reads input file and returns answer
- [x] Write integration test using example-input.txt
- [x] Test against example data (expected: 2 safe reports)
- [x] Test against actual puzzle input

## Data Structures

```go
type Report struct {
    Levels []int
}

func (r Report) IsSafe() bool
func ParseInput(filename string) ([]Report, error)
func CountSafeReports(reports []Report) int
func SolvePart1(filename string) (int, error)
```

## Test Cases from Example
- `7 6 4 2 1` → Safe (decreasing by 1-2)
- `1 2 7 8 9` → Unsafe (increase of 5)
- `9 7 6 2 1` → Unsafe (decrease of 4) 
- `1 3 2 4 5` → Unsafe (mixed directions)
- `8 6 4 4 1` → Unsafe (no change between 4 4)
- `1 3 6 7 9` → Safe (increasing by 1-3)

Expected result: 2 safe reports