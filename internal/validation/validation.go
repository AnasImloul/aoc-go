// Package validation provides input validation utilities for Advent of Code commands.
package validation

import (
	"fmt"
	"time"
)

const (
	// MinYear is the minimum valid Advent of Code year.
	MinYear = 2015
	// MinDay is the minimum valid day.
	MinDay = 1
	// MaxDay is the maximum valid day in Advent of Code.
	MaxDay = 25
	// AdventMonth is December (month 12).
	AdventMonth = 12
)

// GetMaxYear returns the current year as the maximum valid year.
func GetMaxYear() int {
	return time.Now().Year()
}

// ValidateYear checks if a year is within valid Advent of Code range.
// The maximum year is the current year.
func ValidateYear(year int) error {
	maxYear := GetMaxYear()
	if year < MinYear || year > maxYear {
		return fmt.Errorf("year must be between %d and %d, got %d", MinYear, maxYear, year)
	}
	return nil
}

// ValidateDay checks if a day is within valid Advent of Code range.
func ValidateDay(day int) error {
	if day < MinDay || day > MaxDay {
		return fmt.Errorf("day must be between %d and %d, got %d", MinDay, MaxDay, day)
	}
	return nil
}

// ValidateYearAndDay validates both year and day.
// For the current year, it also checks that the day is not in the future.
// Advent of Code problems are released daily from December 1-25.
func ValidateYearAndDay(year, day int) error {
	if err := ValidateYear(year); err != nil {
		return err
	}
	if err := ValidateDay(day); err != nil {
		return err
	}

	// For the current year, check if the day is in the future
	maxYear := GetMaxYear()
	if year == maxYear {
		now := time.Now()
		currentMonth := int(now.Month())
		currentDay := now.Day()

		// If we're not in December yet, no days are available
		if currentMonth < AdventMonth {
			return fmt.Errorf("advent of code %d has not started yet (starts December 1st)", year)
		}

		// If we're in December, check if the requested day is today or in the past
		if currentMonth == AdventMonth {
			if day > currentDay {
				return fmt.Errorf("day %d is not yet available (today is December %d)", day, currentDay)
			}
		}
		// If we're past December, all days 1-25 are valid
	}

	return nil
}

