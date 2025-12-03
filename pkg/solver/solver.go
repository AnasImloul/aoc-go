// Package solver provides the interface and base implementation for Advent of Code solutions.
//
// # Usage
//
// Solutions embed the Base struct and implement FirstPart and SecondPart functions:
//
//	type Day struct {
//	    solver.Base
//	    // parsed data fields
//	}
//
//	func init() {
//	    s := &Day{Base: solver.Base{Year: 2024, Day: 1}}
//	    s.FirstPart = s.firstPart
//	    s.SecondPart = s.secondPart
//	    s.ParseFunc = s.parse
//	    registry.Register(2024, 1, s)
//	}
//
// # Thread Safety
//
// Solver instances are not thread-safe. Each Solve() call may modify the solver's
// internal state (via ParseFunc). For concurrent execution of multiple parts,
// create separate solver instances.
package solver

import (
	"fmt"

	"github.com/AnasImloul/aoc-go/pkg/input"
)

// Solver is the interface that all day solutions must implement.
type Solver interface {
	ReadInput() string
	Solve(part string) any
}

// PartFunc is a function that solves a part of a day's puzzle.
type PartFunc func() any

// ParseFunc is a function that parses the input for a day's puzzle.
type ParseFunc func(string) error

// Base provides a default implementation of the Solver interface.
// Solutions should embed this struct and set the Year, Day, FirstPart, and SecondPart fields.
type Base struct {
	Year       int
	Day        int
	FirstPart  PartFunc
	SecondPart PartFunc
	ParseFunc  ParseFunc
}

// Solve handles executing the correct part logic.
// Panics if the part is not implemented or if parsing fails.
func (b Base) Solve(part string) any {
	// Parse input if ParseFunc is defined
	if b.ParseFunc != nil {
		inputData := b.ReadInput()
		if err := b.ParseFunc(inputData); err != nil {
			panic(fmt.Sprintf("failed to parse input for year %d day %d: %v", b.Year, b.Day, err))
		}
	}

	switch part {
	case "first":
		if b.FirstPart != nil {
			return b.FirstPart()
		}
		panic(fmt.Sprintf("first part not implemented for year %d day %d", b.Year, b.Day))
	case "second":
		if b.SecondPart != nil {
			return b.SecondPart()
		}
		panic(fmt.Sprintf("second part not implemented for year %d day %d", b.Year, b.Day))
	default:
		panic(fmt.Sprintf("invalid part %q: must be 'first' or 'second'", part))
	}
}

// ReadInput reads the input for the given year and day.
func (b Base) ReadInput() string {
	return input.Read(b.Year, b.Day)
}

// ReadLines returns a channel that streams lines from the input file.
func (b Base) ReadLines() <-chan string {
	return input.ReadLines(b.Year, b.Day)
}
