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
	var doStarts []int
	var dontEnds []int

	for _, pos := range doPositions {
		doStarts = append(doStarts, pos[0]) // Get the starting index of `do()`
	}

	for _, pos := range dontPositions {
		dontEnds = append(dontEnds, pos[1]) // Get the ending index of `don't()`
	}

	// Use two-pointer technique to rangeActivatedMemory `do()` with `don't()`
	var rangesActivatedMemory [][]int
	doIdx, dontIdx := 0, 0

	for doIdx < len(doStarts) && dontIdx < len(dontEnds) {
		if doStarts[doIdx] < dontEnds[dontIdx] {
			rangesActivatedMemory = append(rangesActivatedMemory, []int{doStarts[doIdx], dontEnds[dontIdx]})

			for doIdx < len(doStarts) && doStarts[doIdx] < dontEnds[dontIdx] {
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
	for _, rangeActivatedMemory := range rangesActivatedMemory {
		substr := memory[rangeActivatedMemory[0]:rangeActivatedMemory[1]]

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
