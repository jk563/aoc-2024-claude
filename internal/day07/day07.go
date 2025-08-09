package day07

import (
	"bufio"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Equation struct {
	TestValue int
	Operands  []int
}


// Evaluate expression left-to-right with mathematical concatenation
func evaluateExpression(operands []int, operators []string) int {
	if len(operands) == 0 {
		return 0
	}
	if len(operands) == 1 {
		return operands[0]
	}
	
	result := operands[0]
	for i, op := range operators {
		if op == "+" {
			result += operands[i+1]
		} else if op == "*" {
			result *= operands[i+1]
		} else if op == "||" {
			result = concatenateNumbersMath(result, operands[i+1])
		}
	}
	
	return result
}

// Check if equation can be solved with iterator pattern and early termination
func canSolveEquation(testValue int, operands []int, availableOperators []string) bool {
	if len(operands) == 1 {
		return operands[0] == testValue
	}
	
	positions := len(operands) - 1
	operatorCount := len(availableOperators)
	total := 1
	for i := 0; i < positions; i++ {
		total *= operatorCount
	}
	
	// Generate combinations on-demand and test immediately
	for i := 0; i < total; i++ {
		// Generate operators for this combination
		operators := make([]string, positions)
		num := i
		for j := positions - 1; j >= 0; j-- {
			operators[j] = availableOperators[num%operatorCount]
			num /= operatorCount
		}
		
		// Test this combination immediately
		if evaluateExpression(operands, operators) == testValue {
			return true // Early termination!
		}
	}
	
	return false
}

// Parse single equation line
func parseEquation(line string) (Equation, error) {
	parts := strings.Split(line, ": ")
	testValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return Equation{}, err
	}
	
	operandStrs := strings.Fields(parts[1])
	operands := make([]int, len(operandStrs))
	for i, str := range operandStrs {
		operands[i], err = strconv.Atoi(str)
		if err != nil {
			return Equation{}, err
		}
	}
	
	return Equation{TestValue: testValue, Operands: operands}, nil
}

// Parse input file following day01 pattern
func parseInput(filename string) ([]Equation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var equations []Equation
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		equation, err := parseEquation(line)
		if err != nil {
			return nil, err
		}

		equations = append(equations, equation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return equations, nil
}

// Part 1 solution: + and * operators only (parallel)
func SolvePart1(filename string) (int, error) {
	equations, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	
	part1Operators := []string{"+", "*"}
	return solveEquationsParallel(equations, part1Operators), nil
}

// Mathematical concatenation: a * 10^(digits in b) + b
// Uses Go's optimized math.Pow10
func concatenateNumbersMath(a, b int) int {
	// Count digits in b
	digits := 1
	if b != 0 {
		digits = 0
		temp := b
		if temp < 0 {
			temp = -temp
		}
		for temp > 0 {
			digits++
			temp /= 10
		}
	}
	return a*int(math.Pow10(digits)) + b
}

// Part 2 solution: +, *, and || operators (parallel)
func SolvePart2(filename string) (int, error) {
	equations, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	
	part2Operators := []string{"+", "*", "||"}
	return solveEquationsParallel(equations, part2Operators), nil
}

// Worker pool result for parallel processing
type EquationResult struct {
	TestValue int
	Solvable  bool
}

// Process equations in parallel using worker pool
func solveEquationsParallel(equations []Equation, availableOperators []string) int {
	return solveEquationsParallelWithWorkers(equations, availableOperators, runtime.NumCPU())
}

// Process equations in parallel with custom worker count
func solveEquationsParallelWithWorkers(equations []Equation, availableOperators []string, numWorkers int) int {
	equationChan := make(chan Equation, len(equations))
	resultChan := make(chan EquationResult, len(equations))
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for equation := range equationChan {
				solvable := canSolveEquation(equation.TestValue, equation.Operands, availableOperators)
				resultChan <- EquationResult{
					TestValue: equation.TestValue,
					Solvable:  solvable,
				}
			}
		}()
	}
	
	// Send equations to workers
	go func() {
		for _, equation := range equations {
			equationChan <- equation
		}
		close(equationChan)
	}()
	
	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	
	totalCalibrationResult := 0
	for result := range resultChan {
		if result.Solvable {
			totalCalibrationResult += result.TestValue
		}
	}
	
	return totalCalibrationResult
}


