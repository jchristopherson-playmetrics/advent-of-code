package solutions

import "fmt"

// Solution represents a day's puzzle solution
type Solution interface {
	Part1(input string) string
	Part2(input string) string
}

// GetSolution returns the solution for a given year and day
func GetSolution(year, day int) (Solution, error) {
	switch year {
	case 2025:
		return get2025Solution(day)
	default:
		return nil, fmt.Errorf("no solutions found for year %d", year)
	}
}
