package main

import (
	"fmt"
	"os"
	"strconv"
)

type InputArgs struct {
	pathToInputFolder string
	day               int // day of December
	puzzle            int // puzzle of the day
}

func parsePathToInputFolderArg(arg string) (string, error) {
	fileInfo, err := os.Stat(arg)

	if err != nil {
		return "", fmt.Errorf("invalid path to input folder")
	}

	if !fileInfo.IsDir() {
		return "", fmt.Errorf("path to input folder is not a directory")
	}

	return arg, nil
}

func parseDayArg(arg string) (int, error) {
	day, err := strconv.Atoi(arg)

	if err != nil {
		return -1, fmt.Errorf("day must be an integer")
	}

	if !(1 <= day && day <= 24) {
		return -1, fmt.Errorf("day must be between 1 and 25")
	}

	return day, nil
}

func parsePuzzleArg(arg string) (int, error) {
	puzzle, err := strconv.Atoi(arg)

	if err != nil {
		return -1, fmt.Errorf("puzzle must be an integer")
	}

	if !(puzzle == 1 || puzzle == 2) {
		return -1, fmt.Errorf("puzzle must be either 1 or 2")
	}

	return puzzle, nil
}

func parseInputArgs(args []string) (InputArgs, error) {
	args = args[1:]

	if len(args) != 3 {
		return InputArgs{}, fmt.Errorf("expected 3 arguments")
	}

	pathToInputFolder, err := parsePathToInputFolderArg(args[0])

	if err != nil {
		return InputArgs{}, err
	}

	day, err := parseDayArg(args[1])

	if err != nil {
		return InputArgs{}, err
	}

	puzzle, err := parsePuzzleArg(args[2])

	if err != nil {
		return InputArgs{}, err
	}

	return InputArgs{pathToInputFolder, day, puzzle}, nil
}
