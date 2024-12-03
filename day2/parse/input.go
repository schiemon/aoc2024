package parse

type Input struct {
	levels [][]int
}

func NewInput(levels [][]int) *Input {
	return &Input{
		levels: levels,
	}
}

func (i *Input) Level(l int) []int {
	return i.levels[l]
}

func (i *Input) NumLevels() int {
	return len(i.levels)
}
