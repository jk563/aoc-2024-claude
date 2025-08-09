package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Position represents a coordinate on the grid
type Position struct {
	Row, Col int
}

// Direction represents the four cardinal directions
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Guard represents the guard's current state
type Guard struct {
	Position  Position
	Direction Direction
}

// GuardState represents a guard's state (position + direction) for loop detection
type GuardState struct {
	Position  Position
	Direction Direction
}

// Grid represents the lab map
type Grid struct {
	Cells [][]rune
}

func parseInput(filename string) (*Grid, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	cells := make([][]rune, len(lines))
	for i, line := range lines {
		cells[i] = []rune(line)
	}

	return &Grid{Cells: cells}, nil
}

func findGuard(grid *Grid) (*Guard, error) {
	for row := 0; row < len(grid.Cells); row++ {
		for col := 0; col < len(grid.Cells[row]); col++ {
			cell := grid.Cells[row][col]
			var direction Direction
			var found bool

			switch cell {
			case '^':
				direction = Up
				found = true
			case '>':
				direction = Right
				found = true
			case 'v':
				direction = Down
				found = true
			case '<':
				direction = Left
				found = true
			}

			if found {
				return &Guard{
					Position:  Position{Row: row, Col: col},
					Direction: direction,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("guard not found in grid")
}

func isInBounds(grid *Grid, pos Position) bool {
	return pos.Row >= 0 && pos.Row < len(grid.Cells) &&
		pos.Col >= 0 && pos.Col < len(grid.Cells[0])
}

func isObstacle(grid *Grid, pos Position) bool {
	if !isInBounds(grid, pos) {
		return false
	}
	return grid.Cells[pos.Row][pos.Col] == '#'
}

func getNextPosition(pos Position, direction Direction) Position {
	switch direction {
	case Up:
		return Position{pos.Row - 1, pos.Col}
	case Right:
		return Position{pos.Row, pos.Col + 1}
	case Down:
		return Position{pos.Row + 1, pos.Col}
	case Left:
		return Position{pos.Row, pos.Col - 1}
	default:
		return pos
	}
}

func turnRight(direction Direction) Direction {
	switch direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return direction
	}
}

func simulatePatrol(grid *Grid, guard *Guard) int {
	visited := make(map[Position]bool)
	currentGuard := *guard

	for {
		// Mark current position as visited
		visited[currentGuard.Position] = true

		// Get next position in current direction
		nextPos := getNextPosition(currentGuard.Position, currentGuard.Direction)

		// Check if next position is out of bounds (guard leaves the area)
		if !isInBounds(grid, nextPos) {
			break
		}

		// Check if next position has obstacle
		if isObstacle(grid, nextPos) {
			// Turn right and stay at current position
			currentGuard.Direction = turnRight(currentGuard.Direction)
		} else {
			// Move forward to next position
			currentGuard.Position = nextPos
		}
	}

	return len(visited)
}

// SolvePart1 solves part 1 of the Day 6 puzzle by simulating the guard's patrol
// and counting the distinct positions visited before leaving the mapped area.
func SolvePart1(filename string) (int, error) {
	grid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	guard, err := findGuard(grid)
	if err != nil {
		return 0, err
	}

	return simulatePatrol(grid, guard), nil
}

// getPatrolPath simulates the original patrol and returns all positions visited
func getPatrolPath(grid *Grid, guard *Guard) map[Position]bool {
	visited := make(map[Position]bool)
	currentGuard := *guard

	for {
		// Mark current position as visited
		visited[currentGuard.Position] = true

		// Get next position in current direction
		nextPos := getNextPosition(currentGuard.Position, currentGuard.Direction)

		// Check if next position is out of bounds (guard leaves the area)
		if !isInBounds(grid, nextPos) {
			break
		}

		// Check if next position has obstacle
		if isObstacle(grid, nextPos) {
			// Turn right and stay at current position
			currentGuard.Direction = turnRight(currentGuard.Direction)
		} else {
			// Move forward to next position
			currentGuard.Position = nextPos
		}
	}

	return visited
}

// simulatePatrolWithLoopDetection simulates the guard's patrol and returns:
// - true if the guard gets stuck in a loop
// - false if the guard leaves the mapped area
func simulatePatrolWithLoopDetection(grid *Grid, guard *Guard) bool {
	visitedStates := make(map[GuardState]bool)
	currentGuard := *guard
	
	for {
		// Create current state
		state := GuardState{
			Position:  currentGuard.Position,
			Direction: currentGuard.Direction,
		}
		
		// Check if we've seen this state before (loop detected)
		if visitedStates[state] {
			return true
		}
		
		// Mark current state as visited
		visitedStates[state] = true
		
		// Get next position in current direction
		nextPos := getNextPosition(currentGuard.Position, currentGuard.Direction)
		
		// Check if next position is out of bounds (guard leaves the area)
		if !isInBounds(grid, nextPos) {
			return false
		}
		
		// Check if next position has obstacle
		if isObstacle(grid, nextPos) {
			// Turn right and stay at current position
			currentGuard.Direction = turnRight(currentGuard.Direction)
		} else {
			// Move forward to next position
			currentGuard.Position = nextPos
		}
	}
}

// SolvePart2 solves part 2 of the Day 6 puzzle by finding all positions where
// placing a single new obstacle would cause the guard to get stuck in a loop.
func SolvePart2(filename string) (int, error) {
	grid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	guard, err := findGuard(grid)
	if err != nil {
		return 0, err
	}

	// Get all positions visited in the original patrol path
	patrolPath := getPatrolPath(grid, guard)
	guardStartPos := guard.Position
	loopPositions := 0

	// Only try placing obstacles at positions the guard would visit
	for pos := range patrolPath {
		// Skip if position is the guard's starting position
		if pos == guardStartPos {
			continue
		}

		// Skip if position already has an obstacle (shouldn't happen in normal patrol)
		if grid.Cells[pos.Row][pos.Col] == '#' {
			continue
		}

		// Temporarily place obstacle
		originalCell := grid.Cells[pos.Row][pos.Col]
		grid.Cells[pos.Row][pos.Col] = '#'

		// Test if this creates a loop
		if simulatePatrolWithLoopDetection(grid, guard) {
			loopPositions++
		}

		// Restore original cell
		grid.Cells[pos.Row][pos.Col] = originalCell
	}

	return loopPositions, nil
}
