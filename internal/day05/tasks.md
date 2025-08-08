# Day 5: Print Queue - Tasks

## Problem Analysis
- Need to parse page ordering rules (X|Y format)
- Need to parse update sequences (comma-separated page numbers)
- Validate which updates follow the ordering rules
- Find middle page number of correctly-ordered updates
- Sum the middle page numbers

## Implementation Plan

### Task 1: Define data structures
- [ ] Create struct to represent an ordering rule (before, after)
- [ ] Create type for update sequence (slice of page numbers)
- [ ] Create struct for the puzzle input (rules + updates)

### Task 2: Parse input file
- [ ] Function to parse ordering rules from first section
- [ ] Function to parse update sequences from second section  
- [ ] Function to parse complete input file (split by blank line)

### Task 3: Validation logic
- [ ] Function to check if an update follows all applicable ordering rules
- [ ] Function to find middle page number of an update sequence

### Task 4: Solve part 1
- [ ] Filter correctly-ordered updates
- [ ] Extract middle page numbers
- [ ] Sum the middle page numbers
- [ ] Main solver function for part 1

### Task 5: Testing and verification
- [ ] Unit tests for parsing functions
- [ ] Unit tests for validation logic
- [ ] Unit tests for middle page calculation
- [ ] Integration test with example input
- [ ] Test with actual puzzle input

## Key Insights
- Only rules involving pages present in an update apply
- Middle page is at index len(update)/2 for odd-length sequences
- Need to check all pairs in each update against applicable rules

## Part 2 Requirements
- Find incorrectly-ordered updates (those that fail validation)
- Sort them according to the ordering rules
- Find middle page numbers of the sorted updates
- Sum the middle page numbers

### Part 2 Implementation Plan

#### Task 6: Sorting function
- [ ] Implement function to sort an update according to ordering rules
- [ ] Use a custom comparator based on the rules
- [ ] Handle cases where no direct rule exists between two pages

#### Task 7: Solve part 2
- [ ] Filter incorrectly-ordered updates
- [ ] Sort each incorrect update using ordering rules
- [ ] Extract middle page numbers from sorted updates
- [ ] Sum the middle page numbers
- [ ] Main solver function for part 2

#### Task 8: Integration and testing
- [ ] Add Part 2 case to main.go solveDayPart function
- [ ] Test with example data (expected: 123)
- [ ] Test with actual puzzle input

## Part 2 Key Insights
- Need topological sort or custom comparison function
- Can use Go's sort.Slice with custom comparator
- For pages A and B: if rule A|B exists, A comes before B
- Need to handle transitive relationships through the rules

## Updated Part 2 Plan (Refined)

### Task 6: Fix ordering function ✓
- [x] Implement `FixUpdateOrder(update Update, rules []OrderingRule) Update`
- [x] Build rule lookup map for efficiency
- [x] Use sort.Slice with custom comparator based on rules
- [x] Handle pages not covered by any rules

### Task 7: Solve part 2 ✓  
- [x] Filter incorrectly-ordered updates using existing `IsValidUpdate()`
- [x] Fix ordering of each invalid update
- [x] Sum middle page numbers of corrected updates
- [x] Create `SolvePart2()` function

### Task 8: Testing ✓
- [x] Unit tests for `FixUpdateOrder()` with example cases
- [x] Test expected transformations:
  - `75,97,47,61,53` → `97,75,47,61,53` (middle: 47)
  - `61,13,29` → `61,29,13` (middle: 29) 
  - `97,13,75,29,47` → `97,75,47,29,13` (middle: 47)
- [x] Integration test expecting sum of 123