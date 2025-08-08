# Advent of Code 2024 Build Automation

# Default recipe to display available commands
default:
    @just --list

# Build the project
build:
    go build -o advent-of-code-2024 .

# Build with optimization flags for release
# Flags: -s removes symbol table, -w removes DWARF debug info, -trimpath removes file paths
build-release:
    go build -ldflags="-s -w" -trimpath -o advent-of-code-2024 .

# Compare build sizes (regular vs optimized)
build-compare:
    @echo "Building regular binary..."
    @go build -o advent-of-code-2024-regular .
    @echo "Building optimized binary..."
    @go build -ldflags="-s -w" -trimpath -o advent-of-code-2024-optimized .
    @echo ""
    @echo "Size comparison:"
    @ls -lh advent-of-code-2024-regular advent-of-code-2024-optimized | awk '{print $9 ": " $5}'
    @echo ""
    @echo "Size reduction:"
    @stat -f%z advent-of-code-2024-regular advent-of-code-2024-optimized | awk 'NR==1{regular=$1} NR==2{optimized=$1} END{reduction=regular-optimized; percent=(reduction/regular)*100; printf "%.0f bytes (%.1f%%)\n", reduction, percent}'
    @rm -f advent-of-code-2024-regular advent-of-code-2024-optimized

# Run all tests
test:
    go test ./...

# Run tests with verbose output
test-verbose:
    go test -v ./...

# Run tests with coverage report
test-coverage:
    go test -cover ./...

# Run benchmarks
bench:
    go test -bench=. ./...

# Clean build artifacts
clean:
    rm -f advent-of-code-2024 advent-of-code-2024-regular advent-of-code-2024-optimized
    go clean ./...

# Run all implemented puzzles
run: build
    ./advent-of-code-2024

# Run specific day (usage: just run-day 1)
run-day day: build
    ./advent-of-code-2024 -day {{day}}

# Run specific day and part (usage: just run-part 1 2)
run-part day part: build
    ./advent-of-code-2024 -day {{day}} -part {{part}}

# Run with debug output
run-debug: build
    ./advent-of-code-2024 -debug

# Format code
fmt:
    go fmt ./...

# Vet code for potential issues
vet:
    go vet ./...

# Tidy dependencies
tidy:
    go mod tidy

# Run golangci-lint
lint:
    golangci-lint run

# Run all quality checks (format, vet, lint, test)
check: fmt vet lint test