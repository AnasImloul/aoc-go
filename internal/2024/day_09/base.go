package day_09

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  9,
	},
}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) readFileSystem() []int {
	var indices []int
	var id = 0

	for index, sizeStr := range d.ReadInput() {
		size := int(sizeStr - '0')
		if index%2 == 0 {
			for i := 0; i < size; i++ {
				indices = append(indices, id)
			}
			id++
		} else {
			for i := 0; i < size; i++ {
				indices = append(indices, -1)
			}
		}
	}

	return indices
}

func (d Day) computeChecksum(indices []int) int64 {
	var res int64
	for i, fileId := range indices {
		if fileId == -1 {
			continue
		}
		res += int64(fileId * i)
	}
	return res
}
