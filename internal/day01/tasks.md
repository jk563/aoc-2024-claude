# Day 1, Part 1: Historian Hysteria - Implementation Plan

## Problem Summary
- We have two lists of location IDs (numbers)
- Need to pair up smallest from left with smallest from right, second smallest with second smallest, etc.
- Calculate distance (absolute difference) for each pair
- Sum all distances to get total distance

## Algorithm Steps
1. Parse input file containing two columns of numbers
2. Split into two separate lists (left and right)
3. Sort both lists in ascending order
4. Iterate through paired elements and calculate absolute differences
5. Sum all differences to get final result

## Implementation Tasks

### [x] Task 1: Create input parsing functionality
- Write function to read input file
- Parse each line into two integers (left and right list values)
- Return two slices of integers
- Unit test: Test parsing with example data

### [x] Task 2: Create list sorting and pairing logic
- Sort both lists in ascending order
- Validate that both lists have the same length
- Unit test: Test sorting with example data

### [x] Task 3: Implement distance calculation
- Create function to calculate absolute difference between two numbers
- Create function to sum distances for all pairs
- Unit test: Test distance calculation with example pairs

### [x] Task 4: Create main solver function
- Combine parsing, sorting, and calculation
- Return final total distance
- Unit test: Test complete solution with example (expected result: 11)

### [x] Task 5: Integrate with main CLI application
- Create day01 package with exported Solve function
- Update main.go to call day01.Solve() when day=1, part=1
- Handle puzzle-input.txt file input only

### [x] Task 6: Test with example data
- Create example-input.txt with provided test case
- Verify output matches expected result (11) in unit tests
- Debug if necessary

### [x] Task 7: Test with puzzle input
- Get puzzle-input.txt from user
- Run solver and submit answer

## Part 2: Similarity Score

### Problem Summary
- Calculate similarity score instead of total distance
- For each number in left list, count how many times it appears in right list
- Multiply number by its frequency and sum all results

### Algorithm Steps
1. Parse input file (reuse existing parseInput function)
2. Count frequency of each number in right list
3. For each number in left list, multiply by its frequency in right list
4. Sum all similarity scores

### Implementation Tasks

### [x] Task 8: Create frequency counting functionality
- Write function to count occurrences of each number in a list
- Return map[int]int for O(1) lookups
- Unit test: Test frequency counting with example data

### [x] Task 9: Implement similarity score calculation
- Create function to calculate similarity score for two lists
- Use frequency map for efficient lookups
- Unit test: Test with example data (expected result: 31)

### [x] Task 10: Create Part 2 solver function
- Combine parsing and similarity calculation
- Return total similarity score
- Unit test: Test complete solution with example

### [x] Task 11: Update main CLI application
- Add case for day=1, part=2 in solveDayPart function
- Test integration

### [x] Task 12: Test with example and puzzle data
- Verify example gives result of 31
- Run against puzzle input

## Test Data (Example)
```
3   4
4   3
2   5
1   3
3   9
3   3
```
Expected result: 11 (Part 1), 31 (Part 2)

## File Structure
```
day01/
├── tasks.md
├── day01.go
├── day01_test.go
├── example-input.txt
└── puzzle-input.txt
```