# Day 3: Mull It Over - Part 1

## Problem Analysis
- Parse corrupted memory looking for valid `mul(X,Y)` instructions
- X and Y are digit numbers (no specific limit mentioned in problem)
- Invalid patterns should be ignored (e.g., `mul(4*`, `mul(6,9!`, `?(12,34)`, `mul ( 2 , 4 )`)
- Sum all results of valid multiplications

## Example
Input: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
Valid instructions: `mul(2,4)`, `mul(5,5)`, `mul(11,8)`, `mul(8,5)`
Result: 2*4 + 5*5 + 11*8 + 8*5 = 8 + 25 + 88 + 40 = 161

## Implementation Plan

### Step 1: Create regex pattern to match valid mul instructions
- [ ] Pattern: `mul\((\d+),(\d+)\)` (allows any number of digits)
- [ ] Test with valid cases: `mul(2,4)`, `mul(123,456)`, `mul(1,1)`, `mul(12345,67890)`
- [ ] Test with invalid cases: `mul(4*`, `mul(6,9!`, `mul ( 2 , 4 )`

### Step 2: Create function to extract numbers from matched groups
- [ ] Parse the two captured groups as integers
- [ ] Multiply them together
- [ ] Return the result

### Step 3: Create main parsing function
- [ ] Use regex to find all valid mul instructions in input
- [ ] Extract and multiply numbers for each match
- [ ] Sum all results
- [ ] Return total sum

### Step 4: Create file I/O functions
- [ ] Function to read input file
- [ ] Function to solve part 1

### Step 5: Integration
- [x] Export SolvePart1 function for main CLI
- [x] Test with example input
- [x] Test with puzzle input

## Part 2: Conditional Instructions

### Problem Analysis
- Add support for `do()` and `don't()` instructions
- `do()` enables future mul instructions
- `don't()` disables future mul instructions  
- Only most recent do()/don't() applies
- mul instructions are enabled at the beginning
- Need to process instructions in order, maintaining state

### Example
Input: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
- `mul(2,4)` - enabled (8)
- `don't()` - disables future muls
- `mul(5,5)` - disabled (ignored)
- `mul(32,64]` - invalid format anyway
- `mul(11,8)` - disabled (ignored)
- `do()` - enables future muls
- `mul(8,5)` - enabled (40)
Result: 8 + 40 = 48

### Implementation Plan

### Step 6: Parse conditional instructions
- [ ] Pattern to match `do()`: `do\(\)`
- [ ] Pattern to match `don't()`: `don't\(\)`
- [ ] Function to find all instructions (mul, do, don't) with positions

### Step 7: Stateful processing
- [ ] Sort all instructions by position in the input
- [ ] Process in order, tracking enabled/disabled state
- [ ] Only process mul instructions when enabled

### Step 8: Integration
- [ ] Implement SolvePart2 function
- [ ] Add to main CLI
- [ ] Test with example and puzzle input