package year2024

import (
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_01"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_02"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_03"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_04"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_05"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_06"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_07"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_08"
	"github.com/AnasImloul/advent-of-code-golang/internal/2024/day_09"
)

func Run(day int, part string) any {
	switch day {
	case 1:
		return day_01.Solver.Solve(part)
	case 2:
		return day_02.Solver.Solve(part)
	case 3:
		return day_03.Solver.Solve(part)
	case 4:
		return day_04.Solver.Solve(part)
	case 5:
		return day_05.Solver.Solve(part)
	case 6:
		return day_06.Solver.Solve(part)
	case 7:
		return day_07.Solver.Solve(part)
	case 8:
		return day_08.Solver.Solve(part)
	case 9:
		return day_09.Solver.Solve(part)
	default:
		return nil
	}
}
