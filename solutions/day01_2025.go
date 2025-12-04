package solutions

import (
	"strconv"
	"strings"
)

// Day01 represents the solution for 2025 Day 1
type Day01 struct{}

const (
	dialSize  = 100
	startPos  = 50
	targetPos = 0
)

// wrap calculates the final position after moving a delta from a start position on a circular dial of given size.
// It also counts how many times the position passes through zero during the movement.
// Returns:
//   end    - the final position on the dial after applying the delta
//   passes - the number of times the position crosses zero
func wrap(start, delta, size int) (end, passes int) {
	end = (start + delta%size + size) % size

	if delta == 0 {
		return end, 0
	}

	steps := delta
	if steps < 0 {
		steps = -steps
	}

	dir := 1
	if delta < 0 {
		dir = -1
	}

	first := (-dir * start) % size
	if first < 0 {
		first += size
	}
	if first == 0 {
		first = size
	}

	if first > steps {
		return end, 0
	}

	passes = 1 + (steps-first)/size
	return end, passes
}

// solve processes a series of dial movement instructions.
// It simulates moving a dial based on each instruction, starting from startPos.
// If part2 is false, it counts the number of times the dial lands exactly on targetPos (zero).
// If part2 is true, it counts the number of times the dial passes over targetPos (zero) during movements.
func solve(input string, part2 bool) int {
	pos := startPos
	zeroCount := 0

	for _, line := range strings.Fields(input) {
		direction := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		delta := num
		if direction == 'L' {
			delta = -num
		}

		end, passes := wrap(pos, delta, dialSize)
		pos = end

		if part2 {
			zeroCount += passes
		} else {
			if pos == targetPos {
				zeroCount++
			}
		}
	}

	return zeroCount
}

// Part1 solves part 1 of the puzzle
func (d *Day01) Part1(input string) string {
	result := solve(input, false)
	return strconv.Itoa(result) + " is the password"
}

// Part2 solves part 2 of the puzzle
func (d *Day01) Part2(input string) string {
	result := solve(input, true)
	return strconv.Itoa(result) + " is the password"
}
