package day4

import (
	"strconv"
)

func (input *Input) isAtXMASCross(element Element) bool {
	topRightCorner := input.GetElement(element.Position().AddDirection(UpRight))
	bottomRightCorner := input.GetElement(element.Position().AddDirection(DownRight))
	bottomLeftCorner := input.GetElement(element.Position().AddDirection(DownLeft))
	topLeftCorner := input.GetElement(element.Position().AddDirection(UpLeft))

	possibleXMASCorners := []struct {
		element           *Element
		diagonalDirection Direction
	}{
		{topRightCorner, DownLeft},
		{bottomRightCorner, UpLeft},
		{bottomLeftCorner, UpRight},
		{topLeftCorner, DownRight},
	}

	for _, corner := range possibleXMASCorners {
		if corner.element == nil {
			return false
		}
	}

	xmasCorners := 0
	for _, corner := range possibleXMASCorners {
		elementsInDiagonal, ok := input.getElementsInDirection(*corner.element, corner.diagonalDirection, len("MAS"))
		if !ok {
			panic("expected to get all elements in diagonal")
		}

		if concatenateValues(elementsInDiagonal) == "MAS" {
			xmasCorners++

			if xmasCorners >= 2 {
				return true
			}
		}
	}

	return false
}

func (input *Input) numberOfXMASCrosses() int {
	numXMASCrosses := 0

	for element := range input.AllElements {
		if input.isAtXMASCross(element) {
			numXMASCrosses++
		}
	}

	return numXMASCrosses
}

func solvePuzzle2(input *Input) (string, error) {
	numXMASCrosses := input.numberOfXMASCrosses()
	return strconv.Itoa(numXMASCrosses), nil
}

func solvePuzzle2WithRawInput(rawInput []string) (string, error) {
	input, err := Parse(rawInput)

	if err != nil {
		return "", err
	}

	output, err := solvePuzzle2(input)

	if err != nil {
		return "", err
	}

	return output, nil
}
