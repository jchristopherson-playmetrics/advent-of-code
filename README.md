# Advent of Code - Go Solutions

A Golang project to organize and run Advent of Code solutions by year and day.

## Project Structure

```
advent-of-code/
├── main.go                     # Main entry point
├── go.mod                      # Go module file
├── solutions/                  # Solution implementations
│   ├── solution.go            # Solution interface and registry
│   ├── year2025.go            # 2025 solution registry
│   └── day01_2025.go          # 2025 Day 1 solution
└── inputs/                     # Input files organized by year
    └── 2025/
        └── day01.txt          # Input for 2025 Day 1
```

## Usage

Run a specific solution using command-line flags:

```bash
# Run 2025 Day 1 Part 1 (default)
go run main.go -year 2025 -day 1 -part 1

# Run 2025 Day 1 Part 2
go run main.go -year 2025 -day 1 -part 2

# Defaults to current year if -year is omitted
go run main.go -day 1 -part 1
```

### Build and Run

```bash
# Build the executable
go build -o aoc

# Run the executable
./aoc -year 2025 -day 1 -part 1
```

## Adding New Solutions

### 1. Add the input file

Create an input file at `inputs/{year}/day{NN}.txt`:

```bash
# Example for 2025 Day 2
mkdir -p inputs/2025
touch inputs/2025/day02.txt
```

### 2. Implement the solution

Create a new file `solutions/day{NN}_{year}.go`:

```go
package solutions

type Day02 struct{}

func (d *Day02) Part1(input string) string {
    // Implement Part 1 solution
    return "answer"
}

func (d *Day02) Part2(input string) string {
    // Implement Part 2 solution
    return "answer"
}
```

### 3. Register the solution

Add the day to the appropriate year file (e.g., `solutions/year2025.go`):

```go
func get2025Solution(day int) (Solution, error) {
    switch day {
    case 1:
        return &Day01{}, nil
    case 2:
        return &Day02{}, nil  // Add this line
    default:
        return nil, fmt.Errorf("no solution found for 2025 day %d", day)
    }
}
```

### 4. Adding a new year

1. Create a new registry file `solutions/year{YYYY}.go`:

```go
package solutions

import "fmt"

func get2026Solution(day int) (Solution, error) {
    switch day {
    case 1:
        return &Day01_2026{}, nil
    default:
        return nil, fmt.Errorf("no solution found for 2026 day %d", day)
    }
}
```

2. Update `solutions/solution.go` to include the new year:

```go
func GetSolution(year, day int) (Solution, error) {
    switch year {
    case 2025:
        return get2025Solution(day)
    case 2026:
        return get2026Solution(day)  // Add this
    default:
        return nil, fmt.Errorf("no solutions found for year %d", year)
    }
}
```

## Output

The program outputs:
- Year, Day, and Part being solved
- The answer
- Execution time

Example:
```
Year 2025, Day 1, Part 1
Answer: 42
Time: 123.456µs
```

## Requirements

- Go 1.21 or higher

## License

MIT
