package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"advent-of-code-2024/internal/day01"
	"advent-of-code-2024/internal/day02"
	"advent-of-code-2024/internal/day03"
	"advent-of-code-2024/internal/day04"
	"advent-of-code-2024/internal/day05"
	"advent-of-code-2024/internal/day06"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	MinDay  = 1
	MaxDay  = 25
	MinPart = 1
	MaxPart = 2
)

type PuzzleResult struct {
	Day      int
	Part     int
	Result   int
	Duration time.Duration
	Error    error
}

func main() {
	var day = flag.Int("day", 0, fmt.Sprintf("Run specific day (%d-%d)", MinDay, MaxDay))
	var part = flag.Int("part", 0, fmt.Sprintf("Run specific part (%d-%d)", MinPart, MaxPart))
	var debug = flag.Bool("debug", false, "Enable debug mode with detailed output")
	var help = flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	if err := validateArgs(*day, *part); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("\nUse -help for usage information.")
		os.Exit(1)
	}

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	start := time.Now()

	var results []PuzzleResult

	if *day != 0 && *part != 0 {
		results = runSpecificDayPart(*day, *part, *debug)
	} else if *day != 0 {
		results = runSpecificDay(*day, *debug)
	} else {
		results = runAllDays(*debug)
	}

	elapsed := time.Since(start)
	printResultsTable(results, elapsed, *debug)
}

func printResultsTable(results []PuzzleResult, totalTime time.Duration, debug bool) {
	if len(results) == 0 {
		fmt.Println("No puzzles to solve")
		return
	}

	// Calculate column widths
	dayWidth := 3
	partWidth := 4
	resultWidth := 10
	timeWidth := 12
	statusWidth := 6

	for _, r := range results {
		if len(fmt.Sprintf("%d", r.Result)) > resultWidth-2 {
			resultWidth = len(fmt.Sprintf("%d", r.Result)) + 2
		}
		if len(r.Duration.String()) > timeWidth-2 {
			timeWidth = len(r.Duration.String()) + 2
		}
	}

	// Print table header
	fmt.Println(strings.Repeat("─", dayWidth+partWidth+resultWidth+timeWidth+statusWidth+16))
	fmt.Printf("│ %-*s │ %-*s │ %-*s │ %-*s │ %-*s │\n",
		dayWidth, "Day", partWidth, "Part", resultWidth, "Result", timeWidth, "Time", statusWidth, "Status")
	fmt.Println(strings.Repeat("─", dayWidth+partWidth+resultWidth+timeWidth+statusWidth+16))

	// Print results
	for _, r := range results {
		if r.Error != nil {
			fmt.Printf("│ %*d │ %*d │ %-*s │ %-*s │ %-*s │\n",
				dayWidth, r.Day, partWidth, r.Part, resultWidth, "ERROR", timeWidth, r.Duration.String(), statusWidth, "✗")
			if debug {
				fmt.Printf("  Error: %v\n", r.Error)
			}
		} else {
			fmt.Printf("│ %*d │ %*d │ %*d │ %-*s │ %-*s │\n",
				dayWidth, r.Day, partWidth, r.Part, resultWidth, r.Result, timeWidth, r.Duration.String(), statusWidth, "✓")
		}
	}

	fmt.Println(strings.Repeat("─", dayWidth+partWidth+resultWidth+timeWidth+statusWidth+16))

	// Print summary
	solved := 0
	for _, r := range results {
		if r.Error == nil {
			solved++
		}
	}
	fmt.Printf("Summary: %d/%d puzzles solved in %v\n", solved, len(results), totalTime)
}

func validateDay(day int) error {
	if day < MinDay || day > MaxDay {
		return fmt.Errorf("Day must be between %d and %d, got %d", MinDay, MaxDay, day)
	}
	return nil
}

func getInputFilePath(day int) string {
	return filepath.Join("internal", fmt.Sprintf("day%02d", day), "puzzle-input.txt")
}

func validateArgs(day, part int) error {
	// Cannot specify part without day
	if part != 0 && day == 0 {
		return fmt.Errorf("cannot specify part without day")
	}

	// Validate day range if specified
	if day != 0 && (day < MinDay || day > MaxDay) {
		return fmt.Errorf("day must be between %d and %d", MinDay, MaxDay)
	}

	// Validate part range if specified
	if part != 0 && (part < MinPart || part > MaxPart) {
		return fmt.Errorf("part must be %d or %d", MinPart, MaxPart)
	}

	return nil
}

func showHelp() {
	fmt.Println("Advent of Code 2024 Puzzle Solver")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ./advent-of-code-2024 [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Printf("  -day int     Run specific day (%d-%d)\n", MinDay, MaxDay)
	fmt.Printf("  -part int    Run specific part (%d-%d)\n", MinPart, MaxPart)
	fmt.Println("  -debug       Enable debug mode with detailed output")
	fmt.Println("  -help        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./advent-of-code-2024                    # Run all implemented puzzles")
	fmt.Println("  ./advent-of-code-2024 -day 1             # Run both parts of day 1")
	fmt.Println("  ./advent-of-code-2024 -day 1 -part 2     # Run only part 2 of day 1")
	fmt.Println("  ./advent-of-code-2024 -debug             # Run all puzzles with debug output")
}

func runSpecificDayPart(day, part int, debug bool) []PuzzleResult {
	start := time.Now()
	result, err := solveDayPart(day, part)
	duration := time.Since(start)

	return []PuzzleResult{{
		Day:      day,
		Part:     part,
		Result:   result,
		Duration: duration,
		Error:    err,
	}}
}

func runSpecificDay(day int, debug bool) []PuzzleResult {
	var results []PuzzleResult

	// Run part 1
	start1 := time.Now()
	result1, err1 := solveDayPart(day, MinPart)
	duration1 := time.Since(start1)

	results = append(results, PuzzleResult{
		Day:      day,
		Part:     MinPart,
		Result:   result1,
		Duration: duration1,
		Error:    err1,
	})

	// Run part 2
	start2 := time.Now()
	result2, err2 := solveDayPart(day, MaxPart)
	duration2 := time.Since(start2)

	results = append(results, PuzzleResult{
		Day:      day,
		Part:     MaxPart,
		Result:   result2,
		Duration: duration2,
		Error:    err2,
	})

	return results
}

func runAllDays(debug bool) []PuzzleResult {
	var results []PuzzleResult
	for day := MinDay; day <= MaxDay; day++ {
		// Check if input file exists before running any parts for this day
		inputFile := getInputFilePath(day)
		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			continue // Skip this day entirely if no input file exists
		}

		for part := MinPart; part <= MaxPart; part++ {
			start := time.Now()
			result, err := solveDayPart(day, part)
			duration := time.Since(start)

			// Only add results for implemented puzzles (those that don't return "puzzle not implemented")
			if err == nil || (err != nil && err.Error() != "puzzle not implemented") {
				results = append(results, PuzzleResult{
					Day:      day,
					Part:     part,
					Result:   result,
					Duration: duration,
					Error:    err,
				})
			}
		}
	}

	return results
}

func solveDayPart(day, part int) (int, error) {
	inputFile := getInputFilePath(day)

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return 0, fmt.Errorf("input file does not exist: %s", inputFile)
	}

	switch {
	case day == 1 && part == 1:
		return day01.SolvePart1(inputFile)
	case day == 1 && part == 2:
		return day01.SolvePart2(inputFile)
	case day == 2 && part == 1:
		return day02.SolvePart1(inputFile)
	case day == 2 && part == 2:
		return day02.SolvePart2(inputFile)
	case day == 3 && part == 1:
		return day03.SolvePart1(inputFile)
	case day == 3 && part == 2:
		return day03.SolvePart2(inputFile)
	case day == 4 && part == 1:
		return day04.SolvePart1(inputFile)
	case day == 4 && part == 2:
		return day04.SolvePart2(inputFile)
	case day == 5 && part == 1:
		return day05.SolvePart1(inputFile)
	case day == 5 && part == 2:
		return day05.SolvePart2(inputFile)
	case day == 6 && part == 1:
		return day06.SolvePart1(inputFile)
	case day == 6 && part == 2:
		return day06.SolvePart2(inputFile)
	default:
		return 0, fmt.Errorf("puzzle not implemented")
	}
}
