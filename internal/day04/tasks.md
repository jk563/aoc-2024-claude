# Day 4: Ceres Search - Task Plan

## Problem Analysis
- Find all occurrences of the word "XMAS" in a word search grid
- XMAS can appear in 8 directions: horizontal (left/right), vertical (up/down), and diagonal (4 directions)
- Words can be written backwards (so "SAMX" also counts)
- Words can overlap

## Step-by-Step Plan

### 1. Data Structure Setup
- [ ] Create a function to read the grid from input file
- [ ] Represent the grid as a 2D slice of runes or strings
- [ ] Create helper functions for grid dimensions

### 2. Direction Search Implementation  
- [ ] Define all 8 possible directions as coordinate deltas
- [ ] Create a function to check if a word exists starting from a position in a given direction
- [ ] Handle bounds checking to avoid array out of bounds

### 3. Word Search Logic
- [ ] Iterate through each position in the grid
- [ ] For each position, check all 8 directions for "XMAS"
- [ ] Count total occurrences

### 4. Core Functions to Implement
- [ ] `readGrid(filename string) ([][]rune, error)` - read input file into grid
- [ ] `findXMAS(grid [][]rune) int` - main search function
- [ ] `checkDirection(grid [][]rune, row, col, deltaRow, deltaCol int, target string) bool` - check if word exists in specific direction
- [ ] `isValidPosition(grid [][]rune, row, col int) bool` - bounds checking

### 5. Testing Strategy
- [ ] Test with the provided example (should return 18)
- [ ] Test edge cases (word at boundaries, single character grid, etc.)
- [ ] Test individual direction checking

### 6. Integration
- [ ] Create day4.go with exported Solve function
- [ ] Integrate with main CLI to run Day 4 Part 1