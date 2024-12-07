package day4

import (
	"errors"
	"strconv"
)

type Input struct {
	elements [][]Element
}

var (
	Up        = MakeDirection(-1, 0)
	UpRight   = MakeDirection(-1, 1)
	Right     = MakeDirection(0, 1)
	DownRight = MakeDirection(1, 1)
	Down      = MakeDirection(1, 0)
	DownLeft  = MakeDirection(1, -1)
	Left      = MakeDirection(0, -1)
	UpLeft    = MakeDirection(-1, -1)
)

var AllDirections = []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}

type Direction struct {
	rowDelta int
	colDelta int
}

func MakeDirection(rowDelta int, colDelta int) Direction {
	return Direction{rowDelta, colDelta}
}

func (direction Direction) RowDelta() int {
	return direction.rowDelta
}

func (direction Direction) ColDelta() int {
	return direction.colDelta
}

func (direction Direction) Multiply(factor int) Direction {
	return MakeDirection(direction.rowDelta*factor, direction.colDelta*factor)
}

type Position struct {
	row int
	col int
}

func MakePosition(row int, col int) Position {
	return Position{row, col}
}

func (position Position) Row() int {
	return position.row
}

func (position Position) Col() int {
	return position.col
}

func (position Position) AddDirection(direction Direction) Position {
	return MakePosition(position.row+direction.RowDelta(), position.col+direction.ColDelta())
}

type Element struct {
	position Position
	value    string
}

func (element *Element) Equals(other Element) bool {
	return element.position == other.position && element.value == other.value
}

func (element *Element) NotEquals(other Element) bool {
	return !element.Equals(other)
}

func (element *Element) Value() string {
	return element.value
}

func (element *Element) Position() Position {
	return element.position
}

func NewInput(matrix [][]string) *Input {
	err := ValidateInput(matrix)

	input := &Input{
		elements: make([][]Element, len(matrix)),
	}

	for row, rowElements := range matrix {
		for col, elem := range rowElements {
			element := Element{Position{row, col}, elem}
			input.elements[row] = append(input.elements[row], element)
		}
	}

	if err != nil {
		panic(err)
	}

	return input
}

func ValidateInput(matrix [][]string) error {
	for i, row := range matrix {
		for j, elem := range row {
			if len(elem) != 1 {
				return errors.New("matrix element at " + strconv.Itoa(i) + "," + strconv.Itoa(j) + " is not one character")
			}

			if elem != "X" && elem != "M" && elem != "A" && elem != "S" {
				return errors.New("matrix element is not X, M, A or S. Got: '" + elem + "'")
			}
		}
	}

	return nil
}

func (input *Input) GetElement(position Position) *Element {
	if position.row < 0 || position.row >= len(input.elements) {
		return nil
	}

	if position.col < 0 || position.col >= len(input.elements[position.row]) {
		return nil
	}

	return &input.elements[position.row][position.col]
}

func (input *Input) AllElements(yield func(Element) bool) {
	for _, row := range input.elements {
		for _, element := range row {
			yield(element)
		}
	}
}
