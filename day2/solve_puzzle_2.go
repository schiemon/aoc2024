package day2

import (
	"github.com/schiemon/aoc2024/day2/parse"
	"math"
	"slices"
	"strconv"
)

func isValidDiff(diff int) bool {
	if diff == 0 || math.Abs(float64(diff)) > 3 {
		return false
	}

	return true
}

func isSaveIfRemoved(level []int, i int) bool {
	return IsSafe(slices.Concat(level[:i], level[i+1:]))
}

func isSafeWithDeletion(level []int) bool {
	l := len(level)

	if l <= 1 {
		return true
	}

	// Consider removing nothing.
	if IsSafe(level) {
		return true
	}

	// Consider removing the first or second element.
	if isSaveIfRemoved(level, 0) || isSaveIfRemoved(level, 1) {
		return true
	}

	var isIncreasingTotal bool

	// The first and second stay. If their difference is invalid, then there is no way to fix it.
	if !isValidDiff(level[1] - level[0]) {
		return false
	}

	// As the first stay, we now if we are increasing or decreasing.
	if level[1] > level[0] {
		isIncreasingTotal = true
	} else {
		isIncreasingTotal = false
	}

	for i := 2; i < l; i++ {
		diff := level[i] - level[i-1]

		if !isValidDiff(diff) {
			return ((i-1) > 1) && isSaveIfRemoved(level, i-1) || isSaveIfRemoved(level, i)
		}

		if diff > 0 && !isIncreasingTotal || diff < 0 && isIncreasingTotal {
			return ((i-1) > 1) && isSaveIfRemoved(level, i-1) || isSaveIfRemoved(level, i)
		}
	}

	return true
}

func solvePuzzle2(input *parse.Input) (string, error) {
	numValid := 0

	for i := 0; i < input.NumLevels(); i++ {
		level := input.Level(i)
		if isSafeWithDeletion(level) {
			numValid++
		}
	}

	return strconv.Itoa(numValid), nil
}

func solvePuzzle2WithRawInput(rawInput []string) (string, error) {
	input, err := parse.Parse(rawInput)

	if err != nil {
		return "", err
	}

	output, err := solvePuzzle2(input)

	if err != nil {
		return "", err
	}

	return output, nil
}
