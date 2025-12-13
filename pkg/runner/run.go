// Package runner provides utilities for executing Advent of Code solutions.
package runner

import (
	"github.com/AnasImloul/aoc-go/pkg/registry"
)

// Solution runs the solution for the given year, day, and part.
// It retrieves the solver from the registry and executes the specified part.
// Returns nil if no solver is registered for the specified year/day.
// The part parameter should be "first" or "second" (normalized by the caller).
func Solution(year, day int, part string) any {
	solver := registry.Get(year, day)
	if solver == nil {
		return nil
	}
	return solver.Solve(part)
}
