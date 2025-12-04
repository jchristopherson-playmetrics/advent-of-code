package solutions

import "fmt"

func get2025Solution(day int) (Solution, error) {
	switch day {
	case 1:
		return &Day01{}, nil
	case 2:
		return &Day02{}, nil
	default:
		return nil, fmt.Errorf("no solution found for 2025 day %d", day)
	}
}
