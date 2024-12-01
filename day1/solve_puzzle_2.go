package day1

import (
	"github.com/schiemon/aoc2024/day1/parse"
	"strconv"
)

func solvePuzzle2(input *parse.Input) (string, error) {
	firstList := input.FirstList()
	secondList := input.SecondList()

	secondListCounter := make(map[int]int)

	for _, elem := range secondList {
		secondListCounter[elem]++
	}

	totalSimilarityScore := 0

	for i := 0; i < len(firstList); i++ {
		leftElem := firstList[i]
		similarityScore := leftElem * secondListCounter[leftElem]
		totalSimilarityScore += similarityScore
	}

	return strconv.Itoa(totalSimilarityScore), nil
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
