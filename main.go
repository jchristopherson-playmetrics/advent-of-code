package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jchristopherson-playmetrics/advent-of-code/solutions"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "Advent of Code year")
	day := flag.Int("day", 1, "Advent of Code day")
	part := flag.Int("part", 1, "Part number (1 or 2)")
	
	flag.Parse()

	// Get the solution for the specified year and day
	solution, err := solutions.GetSolution(*year, *day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Read input file
	inputPath := fmt.Sprintf("inputs/%d/day%02d.txt", *year, *day)
	input, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	// Run the appropriate part
	var answer string
	start := time.Now()
	
	switch *part {
	case 1:
		answer = solution.Part1(string(input))
	case 2:
		answer = solution.Part2(string(input))
	default:
		fmt.Fprintf(os.Stderr, "Invalid part number: %d (must be 1 or 2)\n", *part)
		os.Exit(1)
	}
	
	elapsed := time.Since(start)

	// Print the result
	fmt.Printf("Year %d, Day %d, Part %d\n", *year, *day, *part)
	fmt.Printf("Answer: %s\n", answer)
	fmt.Printf("Time: %v\n", elapsed)
}
