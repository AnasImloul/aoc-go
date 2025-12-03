package solver

import (
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
func (b Base) Solve(part string) any {
	// Parse input if ParseFunc is defined
	if b.ParseFunc != nil {
		inputData := b.ReadInput()
		if err := b.ParseFunc(inputData); err != nil {
			panic("Failed to parse input: " + err.Error())
		}
	}

	switch part {
	case "first":
		if b.FirstPart != nil {
			return b.FirstPart()
		} else {
			panic("First part not implemented")
		}
	case "second":
		if b.SecondPart != nil {
			return b.SecondPart()
		} else {
			panic("Second part not implemented")
		}
	default:
		panic("Invalid part ('first', 'second')")
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


