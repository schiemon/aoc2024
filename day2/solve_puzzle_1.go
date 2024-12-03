package day2

import (
	"github.com/schiemon/aoc2024/day2/parse"
	"math"
	"strconv"
)

func IsSafe(level []int) bool {
	isIncreasingTotal := false
	isDecreasingTotal := false

	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 {
			if isDecreasingTotal {
				return false
			} else {
				isIncreasingTotal = true
			}
		} else {
			if isIncreasingTotal {
				return false
			} else {
				isDecreasingTotal = true
			}
		}
	}

	return true
}

func solvePuzzle1(input *parse.Input) (string, error) {
	numValid := 0

	for i := 0; i < input.NumLevels(); i++ {
		level := input.Level(i)

		if IsSafe(level) {
			numValid++
		}
	}

	return strconv.Itoa(numValid), nil
}

func solvePuzzle1WithRawInput(rawInput []string) (string, error) {
	input, err := parse.Parse(rawInput)

	if err != nil {
		return "", err
	}

	output, err := solvePuzzle1(input)

	if err != nil {
		return "", err
	}

	return output, nil
}
