package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(rawInput []string) (*Input, error) {
	var firstList []int
	var secondList []int

	for _, line := range rawInput {
		lineParts := strings.Fields(line)

		if len(lineParts) != 2 {
			return nil, fmt.Errorf("each line must contain two elements")
		}

		elemForFirstList, err := strconv.Atoi(lineParts[0])

		if err != nil {
			return nil, err
		}

		elemForSecondList, err := strconv.Atoi(lineParts[1])

		if err != nil {
			return nil, err
		}

		firstList = append(firstList, elemForFirstList)
		secondList = append(secondList, elemForSecondList)
	}

	if len(firstList) != len(secondList) {
		return nil, fmt.Errorf("first list and second list must have the same length")
	}

	return NewInput(firstList, secondList), nil
}
