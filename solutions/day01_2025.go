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

func dontReInventTheWheel(input string, part2 bool) int {
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
	result := dontReInventTheWheel(input, false)
	return strconv.Itoa(result) + " is the password"
}

// Part2 solves part 2 of the puzzle
func (d *Day01) Part2(input string) string {
	result := dontReInventTheWheel(input, true)
	return strconv.Itoa(result) + " is the password"
}
