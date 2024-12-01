package main

import (
	"fmt"
	"os"
)

func getUsage() string {
	return "Usage: go run solve_puzzle_1.go <path to input folder> <day> <puzzle>"
}

// Usage: go run solve_puzzle_1.go <path-to-input-folder> <day> <puzzle>
//
// The input folder must follow this structure:
//
// <path-to-input-folder>/
// ├── day1/
// │   ├── 1.txt
// │   └── 2.txt
// ├── day2/
// │   ├── 1.txt
// │   └── 2.txt
// ...
func main() {
	inputArgs, err := parseInputArgs(os.Args)

	if err != nil {
		// print also usage information with gtUsage()
		fmt.Fprintln(os.Stderr, fmt.Sprintf("error occurred while parsing input arguments: %v", err))
		fmt.Fprintln(os.Stderr, getUsage())
		os.Exit(1)
	}

	config, err := NewConfigFromInputArgs(inputArgs)

	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("error occurred while building configuration: %v", err))
		os.Exit(1)
	}

	solution, err := run(config)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(solution)
}
