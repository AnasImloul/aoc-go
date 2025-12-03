package runner

import (
	"github.com/AnasImloul/aoc-go/pkg/registry"
)

// Solution runs the solution for the given year, day, and part.
// Returns nil if no solver is registered for the specified year/day.
func Solution(year, day int, part string) any {
	solver := registry.Get(year, day)
	if solver == nil {
		return nil
	}
	return solver.Solve(part)
}


