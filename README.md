# Advent of Code Solutions

A Golang project to organize and run Advent of Code solutions by year and day.

## Project Structure

```text
advent-of-code/
├── main.go                 # Main entry point with CLI
├── go.mod                  # Go module configuration
├── solutions/
│   ├── solution.go         # Solution interface
│   ├── year2025.go         # 2025 solution registry
│   ├── day01_2025.go       # 2025 Day 1 solution
│   └── day02_2025.go       # 2025 Day 2 solution
└── inputs/
    └── 2025/
        ├── day01.txt
        └── day02.txt
```

## Adding New Days

The solution registry uses an auto-registration pattern. When you create a new day solution, it automatically registers itself.

### Steps:

1. **Create a new file** `dayXX_2025.go` in the `solutions/` directory

2. **Implement the day struct** and register it:

```go
package solutions

// Day03 represents the solution for 2025 Day 3
type Day03 struct{}

func init() {
    register2025(3, func() Solution { return &Day03{} })
}

// Part1 solves part 1 of the puzzle
func (d *Day03) Part1(input string) string {
    // Implementation here
    return "answer"
}

// Part2 solves part 2 of the puzzle
func (d *Day03) Part2(input string) string {
    // Implementation here
    return "answer"
}
```

3. **Add the input file** at `inputs/2025/day03.txt`

4. **Run it**:
```bash
go run main.go -year 2025 -day 3 -part 1
```

## The Registration System

Each day calls `register2025(day, factory)` in its `init()` function. This registers a factory function that creates a new instance of that day's solution when needed. The registry automatically handles lookups by day number.

**No need to manually update a registry file** — just create your day solution and it registers itself!

## Building & Running

```bash
# Run with defaults (2025, Day 1, Part 1)
go run main.go

# Run a specific day and part
go run main.go -year 2025 -day 5 -part 2

# Build the project
go build -o aoc

# Run the built binary
./aoc -year 2025 -day 15 -part 1
```

## Debugging with Delve

Press `F5` in VS Code to start debugging with one of the pre-configured launch configurations.
