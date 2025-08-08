package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"advent-of-code-2024/internal/day01"
	"advent-of-code-2024/internal/day02"
	"advent-of-code-2024/internal/day03"
	"advent-of-code-2024/internal/day04"
	"advent-of-code-2024/internal/day05"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	MinDay  = 1
	MaxDay  = 25
	MinPart = 1
	MaxPart = 2
)

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

	if *day != 0 && *part != 0 {
		runSpecificDayPart(*day, *part, *debug)
	} else if *day != 0 {
		runSpecificDay(*day, *debug)
	} else {
		runAllDays(*debug)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total execution time: %v\n", elapsed)
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

func runSpecificDayPart(day, part int, debug bool) {
	if debug {
		log.Debug().Int("day", day).Int("part", part).Msg("Running specific day and part")
	}

	result, err := solveDayPart(day, part)
	if err != nil {
		fmt.Printf("Day %02d Part %d: Error - %v\n", day, part, err)
		return
	}

	fmt.Printf("Day %02d Part %d: %d\n", day, part, result)
}

func runSpecificDay(day int, debug bool) {
	if debug {
		log.Debug().Int("day", day).Msg("Running specific day")
	}

	result1, err1 := solveDayPart(day, MinPart)
	if err1 != nil {
		fmt.Printf("Day %02d Part 1: Error - %v\n", day, err1)
	} else {
		fmt.Printf("Day %02d Part 1: %d\n", day, result1)
	}

	result2, err2 := solveDayPart(day, MaxPart)
	if err2 != nil {
		fmt.Printf("Day %02d Part 2: Error - %v\n", day, err2)
	} else {
		fmt.Printf("Day %02d Part 2: %d\n", day, result2)
	}
}

func runAllDays(debug bool) {
	if debug {
		log.Debug().Msg("Running all implemented days")
	}

	puzzlesSolved := 0
	for day := MinDay; day <= MaxDay; day++ {
		for part := MinPart; part <= MaxPart; part++ {
			result, err := solveDayPart(day, part)
			if err == nil {
				fmt.Printf("Day %02d Part %d: %d\n", day, part, result)
				puzzlesSolved++
			}
		}
	}

	if puzzlesSolved == 0 {
		fmt.Println("No puzzles to solve")
	}
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
	default:
		return 0, fmt.Errorf("puzzle not implemented")
	}
}
