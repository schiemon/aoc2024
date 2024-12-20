package main

import (
	"testing"
)

const pathToInputFolder = "inputs"

func runPuzzle(day int, puzzle int) (string, error) {
	pathToInputFile, err := GetPathToInputFile(pathToInputFolder, day, puzzle)

	if err != nil {
		return "", err
	}

	testConfig := NewConfig(day, puzzle, pathToInputFile)
	return run(testConfig)
}

func runAndTestPuzzle(day int, puzzle int, expectedOutput string, t *testing.T) {
	output, err := runPuzzle(day, puzzle)

	if err != nil {
		t.Errorf("error occurred while running puzzle: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("expected %v but got %v", expectedOutput, output)
	}
}

func TestRunDay1Puzzle1(t *testing.T) {
	runAndTestPuzzle(1, 1, "3574690", t)
}

func TestRunDay1Puzzle2(t *testing.T) {
	runAndTestPuzzle(1, 2, "22565391", t)
}

func TestRunDay2Puzzle1(t *testing.T) {
	runAndTestPuzzle(2, 1, "585", t)
}

func TestRunDay2Puzzle2(t *testing.T) {
	runAndTestPuzzle(2, 2, "626", t)
}

func TestRunDay3Puzzle1(t *testing.T) {
	runAndTestPuzzle(3, 1, "168539636", t)
}

func TestRunDay3Puzzle2(t *testing.T) {
	runAndTestPuzzle(3, 2, "97529391", t)
}

func TestRunDay4Puzzle1(t *testing.T) {
	runAndTestPuzzle(4, 1, "2462", t)
}

func TestRunDay4Puzzle2(t *testing.T) {
	runAndTestPuzzle(4, 2, "1877", t)
}
