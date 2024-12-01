package day1

import (
	"github.com/schiemon/aoc2024/day1/parse"
	"math"
	"sort"
	"strconv"
)

func solvePuzzle1(input *parse.Input) (string, error) {
	firstList := input.FirstList()
	secondList := input.SecondList()

	sort.Ints(firstList)
	sort.Ints(secondList)

	var diffSum int = 0

	for i := 0; i < len(firstList); i++ {
		diff := int(math.Abs(float64(firstList[i] - secondList[i])))
		diffSum += diff
	}

	return strconv.Itoa(diffSum), nil
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
