package day4

import "strings"

func Parse(rawInput []string) (*Input, error) {
	var matrix [][]string
	for _, rowElements := range rawInput {
		matrix = append(matrix, strings.Split(rowElements, ""))
	}

	var err = ValidateInput(matrix)

	if err != nil {
		return nil, err
	}

	return NewInput(matrix), nil
}
