package solutions

import (
	"strconv"
	"strings"
)

// Day03 represents the solution for 2025 Day 3
type Day03 struct{}

func init() {
    register2025(3, func() Solution { return &Day03{} })
}

func maxBankJoltage(line string, k int) int64 {
    digits := strings.Split(line, "")
    n := len(digits)

    result := make([]string, 0, k)
    start := 0

    for pos := 0; pos < k; pos++ {
        // How many digits we still need *after* this pick
        remain := k - pos - 1

        // The last index we are allowed to pick from (inclusive)
        lastAllowed := n - remain - 1

        maxDigit := "-1"
        maxIdx := start

        // Search inclusive range: start..lastAllowed
        for i := start; i <= lastAllowed; i++ {
            if digits[i] > maxDigit {
                maxDigit = digits[i]
                maxIdx = i
            }
        }

        result = append(result, maxDigit)
        start = maxIdx + 1
    }

    // Convert selected digits to a number
    var val int64
    for _, d := range result {
        n, _ := strconv.Atoi(d)
        val = val*10 + int64(n)
    }

    return val
}




func solveDay3(input string, part2 bool) int {
    k := 2
    if part2 {
        k = 12
    }

    sum := int64(0)
    for _, line := range strings.Fields(input) {
        sum += maxBankJoltage(line, k)
    }
    return int(sum)
}

// Part1 solves part 1 of the puzzle
func (d *Day03) Part1(input string) string {
	result := solveDay3(input, false)
	return strconv.Itoa(result)
}

// Part2 solves part 2 of the puzzle
func (d *Day03) Part2(input string) string {
	result := solveDay3(input, true)
	return strconv.Itoa(result)
}
