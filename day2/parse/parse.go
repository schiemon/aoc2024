package parse

import (
	"strconv"
	"strings"
)

func Parse(rawInput []string) (*Input, error) {
	var levels [][]int
	for _, line := range rawInput {
		var level []int
		lineParts := strings.Fields(line)

		for _, part := range lineParts {
			elem, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			level = append(level, elem)
		}

		levels = append(levels, level)
	}

	return NewInput(levels), nil
}
