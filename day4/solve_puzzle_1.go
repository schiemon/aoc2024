package day4

import (
	"strconv"
)

func (input *Input) getElementsInDirection(element Element, direction Direction, n int) ([]Element, bool) {
	elements := make([]Element, n)
	origin := element.Position()

	for factor := 0; factor < n; factor++ {
		newPosition := origin.AddDirection(direction.Multiply(factor))
		newElement := input.GetElement(newPosition)

		if newElement == nil {
			return nil, false
		}

		elements = append(elements, *newElement)
	}

	return elements, true
}

func (input *Input) numberOfXMASStartingAt(element Element) int {
	numberOfXMAS := 0

	for _, direction := range AllDirections {
		elements, ok := input.getElementsInDirection(element, direction, len("XMAS"))

		if !ok {
			continue
		}

		if concatenateValues(elements) == "XMAS" {
			numberOfXMAS++
		}
	}

	return numberOfXMAS
}

func (input *Input) numberXMASStarting() int {
	numXMAS := 0

	for element := range input.AllElements {
		numXMAS += input.numberOfXMASStartingAt(element)
	}

	return numXMAS
}

func concatenateValues(elements []Element) string {
	var s string

	for _, element := range elements {
		s += element.Value()
	}

	return s
}

func solvePuzzle1(input *Input) (string, error) {
	numXMASStarting := input.numberXMASStarting()
	return strconv.Itoa(numXMASStarting), nil
}

func solvePuzzle1WithRawInput(rawInput []string) (string, error) {
	input, err := Parse(rawInput)

	if err != nil {
		return "", err
	}

	output, err := solvePuzzle1(input)

	if err != nil {
		return "", err
	}

	return output, nil
}
