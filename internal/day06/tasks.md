# Day 6: Guard Gallivant - Tasks

## Problem Summary
Simulate a guard's patrol following specific rules:
1. If obstacle ahead, turn right 90 degrees
2. Otherwise, step forward
3. Count distinct positions visited before leaving the mapped area

## Implementation Plan

### Phase 1: Core Data Structures
- [ ] Define Position struct (row, col)
- [ ] Define Direction enum/constants (Up, Right, Down, Left)
- [ ] Define Guard struct (position, direction)
- [ ] Define Grid struct to represent the map

### Phase 2: Grid Parsing
- [ ] Function to parse input into Grid
- [ ] Function to find guard's starting position and direction
- [ ] Function to check if position is within grid bounds
- [ ] Function to check if position contains obstacle

### Phase 3: Movement Logic
- [ ] Function to get next position in current direction
- [ ] Function to turn right (change direction)
- [ ] Function to simulate one step of patrol
- [ ] Function to check if guard is still in bounds

### Phase 4: Simulation
- [ ] Function to simulate full patrol until guard leaves area
- [ ] Track visited positions using a set/map
- [ ] Return count of distinct positions visited

### Phase 5: Integration
- [ ] Implement SolvePart1 function
- [ ] Add proper error handling
- [ ] Ensure it integrates with main CLI

## Test Cases
- [ ] Test parsing grid and finding guard position
- [ ] Test movement in all four directions
- [ ] Test turning right from each direction
- [ ] Test obstacle collision and turning
- [ ] Test full simulation with example input
- [ ] Test edge cases (guard starting at edge, immediate obstacles)

## Files to Create
- `day06.go` - Main implementation
- `day06_test.go` - Unit tests
- `example-input.txt` - Example from problem (after implementation)
- `puzzle-input.txt` - Actual puzzle input (after testing)