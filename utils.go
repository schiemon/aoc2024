package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GetPathToInputFile(pathToInputFolder string, day int, puzzle int) (string, error) {
	pathToInputFile := filepath.Join(pathToInputFolder, fmt.Sprintf("day%d", day), fmt.Sprintf("%d.txt", puzzle))

	_, err := os.Stat(pathToInputFile)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("input file for day %d puzzle %d does not exist. Looked for file at %s", day, puzzle, pathToInputFile))
	}

	return pathToInputFile, nil
}

func ReadInputFile(pathToInputFile string) ([]string, error) {
	file, err := os.Open(pathToInputFile)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var lines []string
	var scannerErr error

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	scannerErr = scanner.Err()

	return lines, scannerErr
}
