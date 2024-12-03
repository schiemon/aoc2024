package day3

import (
	"regexp"
	"strconv"
)

func solvePuzzle2(input []string) (string, error) {
	// Concatenate all lines into one string
	memory := ""

	for _, line := range input {
		memory += line
		memory += "\n"
	}

	memory = "do()" + memory + "don't()"

	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	doPositions := reDo.FindAllStringIndex(memory, -1)
	dontPositions := reDont.FindAllStringIndex(memory, -1)

	// Extract only the start indices for `do()` and end indices for `don't()`
	var starts []int
	var ends []int

	for _, pos := range doPositions {
		starts = append(starts, pos[0]) // Get the starting index of `do()`
	}

	for _, pos := range dontPositions {
		ends = append(ends, pos[1]) // Get the ending index of `don't()`
	}

	// Use two-pointer technique to match `do()` with `don't()`
	var matches [][]int
	doIdx, dontIdx := 0, 0

	for doIdx < len(starts) && dontIdx < len(ends) {
		if starts[doIdx] < ends[dontIdx] {
			// A valid match
			matches = append(matches, []int{starts[doIdx], ends[dontIdx]})

			for doIdx < len(starts) && starts[doIdx] < ends[dontIdx] {
				// Move the `do()` pointer forward to see when the memory is activated again
				doIdx++
			}

			dontIdx++
		} else {
			// Move the `don't()` pointer forward to find a valid do()-don't() pair
			dontIdx++
		}
	}

	totalResult := int64(0)
	for _, match := range matches {
		substr := memory[match[0]:match[1]]

		// Parse the substring
		result, err := SolvePuzzle1OneLine(substr)

		if err != nil {
			return "", err
		}

		totalResult += result
	}

	return strconv.FormatInt(totalResult, 10), nil
}

func solvePuzzle2WithRawInput(input []string) (string, error) {
	return solvePuzzle2(input)
}
