package solutions

import "fmt"

type SolutionFactory func() Solution

var day2025Registry = map[int]SolutionFactory{}

func register2025(day int, factory SolutionFactory) {
	if _, exists := day2025Registry[day]; exists {
		panic(fmt.Sprintf("duplicate registration for 2025 day %d", day))
	}
	day2025Registry[day] = factory
}

func get2025Solution(day int) (Solution, error) {
	factory, ok := day2025Registry[day]
	if !ok {
		return nil, fmt.Errorf("no solution found for 2025 day %d", day)
	}
	return factory(), nil
}
