package solutions

import (
	"strconv"
	"strings"
)

// Day02 represents the solution for 2025 Day 2
type Day02 struct{}

func RangeSlice(r string) []int {
    p := strings.Split(strings.TrimSpace(r), "-")
    if len(p) != 2 {
        // Malformed input, return empty slice
        return []int{}
    }
    s, err1 := strconv.Atoi(p[0])
    e, err2 := strconv.Atoi(p[1])
    if err1 != nil || err2 != nil || e < s {
        // Malformed numbers or invalid range, return empty slice
        return []int{}
    }
    out := make([]int, e-s+1)
    for i := range out {
        out[i] = s + i
    }
    return out
}

func IsDoubledNumber(n string) bool {
    if len(n)%2 != 0 {
        return false
    }

    half := len(n) / 2
    return n[:half] == n[half:]
}

func HasRepeatedPattern(s string) bool {
    if len(s) < 2 {
        return false
    }
    t := s + s

    // slice off first and last char to avoid self-match
    return strings.Contains(t[1:len(t)-1], s)
}

func solveDay2(input string, part2 bool) int {
 	// pseudocode
	sum := 0
	// with input, split on comma
	parts := strings.Split(input, ",")
	// for each set in the input,
	for _, part := range parts {
		// find all the numbers in the set
		ids := RangeSlice(part)
		
		for _, id := range ids {
			s := strconv.Itoa(id)
			// determine which are invalid (repeat the same numbers twice)
			if !part2 && IsDoubledNumber(s) {
				sum += id
				
			// for part 2, determine if the ids have a repeated pattern
			} else if part2 && HasRepeatedPattern(s) {
				sum += id
			}
		}
	}
	return sum
}

// Part1 solves part 1 of the puzzle
func (d *Day02) Part1(input string) string {
	result := solveDay2(input, false)
	return strconv.Itoa(result)
}

// Part2 solves part 2 of the puzzle
func (d *Day02) Part2(input string) string {
	result := solveDay2(input, true)
	return strconv.Itoa(result)
}
