package main

import (
	"fmt"
	"github.com/schiemon/aoc2024/day1"
	"github.com/schiemon/aoc2024/day2"
	"github.com/schiemon/aoc2024/day3"
)

func getSolver(day int, puzzle int) (func([]string) (string, error), error) {
	var solverForDay func([]string, int) (string, error)

	switch day {
	case 1:
		solverForDay = day1.Solve
	case 2:
		solverForDay = day2.Solve
	case 3:
		solverForDay = day3.Solve
	default:
		return nil, fmt.Errorf("invalid day")
	}

	solverForPuzzleOnDay := func(rawInput []string) (string, error) {
		return solverForDay(rawInput, puzzle)
	}

	return solverForPuzzleOnDay, nil
}

func run(config *Config) (string, error) {
	solver, err := getSolver(config.Day(), config.Puzzle())

	if err != nil {
		return "", fmt.Errorf("error occurred while getting solver: %v", err)
	}

	rawInput, err := ReadInputFile(config.PathToInputFile())

	if err != nil {
		return "", fmt.Errorf("error occurred while reading input file: %v", err)
	}

	output, err := solver(rawInput)

	if err != nil {
		return "", fmt.Errorf("error occurred while solving: %v", err)
	}

	return output, nil
}
