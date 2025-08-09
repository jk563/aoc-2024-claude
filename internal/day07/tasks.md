# Day 7: Bridge Repair - Tasks

## Problem Analysis
- Parse calibration equations with test values and operands
- Insert operators (+, *) between operands to match test values
- Operators are evaluated left-to-right (no precedence rules)
- Find sum of test values from equations that can be made true

## Step-by-step Implementation Plan

### Phase 1: Data Structure and Parsing
- [ ] 1.1 Define Equation struct to hold test value and operands
- [ ] 1.2 Write unit tests for parsing single equation line
- [ ] 1.3 Implement parseEquation function 
- [ ] 1.4 Write unit tests for parsing multiple equations
- [ ] 1.5 Implement parseInput function

### Phase 2: Operator Combination Generation
- [ ] 2.1 Write unit tests for generating all operator combinations for n positions
- [ ] 2.2 Implement generateOperatorCombinations function (using binary representation)
- [ ] 2.3 Write unit tests for edge cases (1 operand, 2 operands)

### Phase 3: Expression Evaluation
- [ ] 3.1 Write unit tests for evaluating expression left-to-right
- [ ] 3.2 Implement evaluateExpression function
- [ ] 3.3 Test with examples from problem description

### Phase 4: Equation Validation
- [ ] 4.1 Write unit tests for checking if equation can be made valid
- [ ] 4.2 Implement canSolveEquation function (tries all operator combinations)
- [ ] 4.3 Test with example equations from problem

### Phase 5: Main Solution
- [ ] 5.1 Write unit tests for calculating total calibration result
- [ ] 5.2 Implement SolvePart1 function
- [ ] 5.3 Test with complete example from problem description

### Phase 6: Integration
- [ ] 6.1 Add Day07 to main.go dispatcher
- [ ] 6.2 Test with example input file
- [ ] 6.3 Test with actual puzzle input

## Key Implementation Details
- Use recursive or iterative approach to generate all 2^(n-1) operator combinations
- Evaluate expressions strictly left-to-right without operator precedence
- Sum only the test values from equations that can be solved

## Part 2 Requirements
- Add concatenation operator (||) that combines digits: 12 || 345 = 12345
- Now have 3 operators: +, *, ||
- Need to generate 3^(n-1) combinations instead of 2^(n-1)
- Examples that work with concatenation:
  - 156: 15 6 → 15 || 6 = 156
  - 7290: 6 8 6 15 → 6 * 8 || 6 * 15 = 48 || 90 = 4890 (incorrect example in problem)
  - 192: 17 8 14 → 17 || 8 + 14 = 178 + 14 = 192
- Total for all solvable equations should be 11387