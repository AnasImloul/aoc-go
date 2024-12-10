package day_10

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/day"
	"strconv"
)

type Day struct {
	day.Base
}

var Solver = Day{
	Base: day.Base{
		Year: 2024,
		Day:  10,
	},
}

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func init() {
	Solver.FirstPart = Solver.firstPart
	Solver.SecondPart = Solver.secondPart
}

func (d Day) readGrid() [][]int {
	grid := make([][]int, 0)

	for line := range d.ReadLines() {
		grid = append(grid, make([]int, len(line)))
		for i, r := range line {
			height, _ := strconv.Atoi(string(r))
			grid[len(grid)-1][i] = height
		}
	}

	return grid
}
