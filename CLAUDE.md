# Project Background

- This project is orientated around solving coding puzzles from Advent of Code
- There is one puzzle per day for each day from Decmber 1st to December 25.
- Each day's puzzle comprises two parts which both need to be solve
- Part 1 of each day needs to be solved before the instructions for part 2 are available
- The project will have one primary executable stored at the root of the directory, with a subdirectory module for each day

# Project Technology

- Golang using stdlib
- The only external dependency should be zerolog which can be used for debugging and logging

# Core CLI features
- Allows running against a specific day, specific part of a day, or all existing day
- The main executable will start a timer, execute all implemented puzzles, and return the total time taken and puzzles solved count
- In debug mode, it will print each day, part, and puzzle outcome for each part
- It will provide a helpful help message detailing the options available

# Core CLI Workflow
- For each task, we will think hard to create a step by step plan for how we will solve it
    - This plan will be a tasks.md file in that directory
    - The plan will be broken down into small steps that can be checked off as they are completed
    - We will only move on to the implementation phase after explicit approval of the plan

# Puzzle Development workflow

- For the root binary development, implement in the root of the directory
- For each day, a new submodule will be created. Existing days can be reused for part 2 implementations
- We will implement a single part at a time
- Once a plan is approved, we will implement each uncompleted sub task one by one
    - For each subtask, identify appropriate unit tests for that task and write them first as we follow test driven development
    - When a test is written, run it to ensure it fails.
    - When we have a failing test, implement the minimum functionality to get the test passing
    - When the task is complete, check it off
    - Once all tests are passing, move on to the next sub task until all tasks are complete
- Once we have all tests passing, ask for an example-input.txt file to be created in the directory and the expected answer
    - Run the code against the example input file to verify correctness
    - If it fails, identify bugs and create tasks for them
    - If it passs, ask for puzzle-input.txt to be created in the directory
    - Run the CLI for that day and part, if it fails identify bugs and add to tasks
    - If it succeeds, ask me if the answer is correct, if it isn't review the code and tests for errors
- Review the code for opportunities to refactor or otherwise improve the code. List them out as a numbered list and ask which ones are worth doing.

# Implementation Notes

- Part 2 of each day may require refactoring part 1, and / or creating reusable functions that both use
- Do NOT use reflection
- Only add imports as they are used
- ALWAYS write tests first for TDD
- NO tests are required for data structures

# Comman Commands (justfile)

 Prefer using these just commands instead of executing commands directly. The project includes a justfile for common development tasks. Use `just <command>` to run these:

## Build Commands
- `build` - Compile the project to `advent-of-code-2024` executable
- `clean` - Remove build artifacts and clean Go cache

## Testing Commands  
- `test` - Run all tests
- `test-verbose` - Run tests with verbose output
- `test-coverage` - Run tests and show coverage report
- `bench` - Run benchmark tests

## Running Commands
- `run` - Build and run all implemented puzzles
- `run-day <day>` - Build and run both parts of a specific day (e.g., `just run-day 1`)
- `run-part <day> <part>` - Build and run a specific day and part (e.g., `just run-part 1 2`)
- `run-debug` - Build and run with debug output enabled

## Code Quality Commands
- `fmt` - Format all Go code
- `vet` - Run Go vet for potential issues  
- `tidy` - Clean up module dependencies
- `check` - Run format, vet, and test in sequence

## Usage
Run `just` without arguments to see all available commands.

