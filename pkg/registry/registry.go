package registry

import (
	"github.com/AnasImloul/aoc-go/pkg/solver"
)

var solvers = make(map[int]map[int]solver.Solver)

// Register adds a solver to the registry for the given year and day.
// This should be called from the init() function of each solution.
func Register(year, day int, s solver.Solver) {
	if solvers[year] == nil {
		solvers[year] = make(map[int]solver.Solver)
	}
	solvers[year][day] = s
}

// Get retrieves a solver for the given year and day.
// Returns nil if no solver is registered for the specified year/day.
func Get(year, day int) solver.Solver {
	if yearSolvers, ok := solvers[year]; ok {
		return yearSolvers[day]
	}
	return nil
}

// GetYears returns all years that have registered solvers.
func GetYears() []int {
	years := make([]int, 0, len(solvers))
	for year := range solvers {
		years = append(years, year)
	}
	return years
}

// GetDays returns all days that have registered solvers for a given year.
func GetDays(year int) []int {
	if yearSolvers, ok := solvers[year]; ok {
		days := make([]int, 0, len(yearSolvers))
		for day := range yearSolvers {
			days = append(days, day)
		}
		return days
	}
	return nil
}


