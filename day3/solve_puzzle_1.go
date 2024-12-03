package day3

import (
	"regexp"
	"strconv"
)

func SolvePuzzle1OneLine(input string) (int64, error) {
	re := regexp.MustCompile(`(?m)mul\((:?\d{1,3}),(:?\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var pairs [][2]int

	for _, match := range matches {
		if len(match) > 2 {
			firstFactor, err := strconv.Atoi(match[1])

			if err != nil {
				return -1, err
			}

			secondFactor, err := strconv.Atoi(match[2])

			if err != nil {
				return -1, err
			}

			pairs = append(pairs, [2]int{firstFactor, secondFactor})
		}
	}

	sumOfProducts := int64(0)

	for _, pair := range pairs {
		sumOfProducts += int64(pair[0] * pair[1])
	}

	return sumOfProducts, nil
}

func solvePuzzle1(input []string) (string, error) {
	result := int64(0)

	for _, line := range input {
		resultLine, err := SolvePuzzle1OneLine(line)

		if err != nil {
			return "", err
		}

		result += resultLine
	}

	return strconv.FormatInt(result, 10), nil
}

func solvePuzzle1WithRawInput(input []string) (string, error) {
	return solvePuzzle1(input)
}
