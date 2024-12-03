package day3

import "fmt"

func Solve(rawInput []string, puzzle int) (string, error) {
	switch puzzle {
	case 1:
		return solvePuzzle1WithRawInput(rawInput)
	case 2:
		return solvePuzzle2WithRawInput(rawInput)
	default:
		return "", fmt.Errorf("invalid puzzle number")
	}
}
