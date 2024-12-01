package parse

type Input struct {
	firstList  []int
	secondList []int
}

func NewInput(firstList []int, secondList []int) *Input {
	if len(firstList) != len(secondList) {
		panic("first list and second list must have the same length")
	}

	return &Input{
		firstList:  firstList,
		secondList: secondList,
	}
}

func (i *Input) FirstList() []int {
	return i.firstList
}

func (i *Input) SecondList() []int {
	return i.secondList
}
