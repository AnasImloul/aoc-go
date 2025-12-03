package registry

import (
	"testing"

	"github.com/AnasImloul/aoc-go/pkg/solver"
)

// mockSolver is a minimal implementation of solver.Solver for testing
type mockSolver struct {
	solver.Base
}

func TestRegisterAndGet(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	s := &mockSolver{
		Base: solver.Base{Year: 2024, Day: 1},
	}

	Register(2024, 1, s)

	got := Get(2024, 1)
	if got != s {
		t.Errorf("Get(2024, 1) = %v, want %v", got, s)
	}
}

func TestGetNonExistent(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	got := Get(2024, 1)
	if got != nil {
		t.Errorf("Get(2024, 1) = %v, want nil", got)
	}
}

func TestGetNonExistentDay(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	s := &mockSolver{
		Base: solver.Base{Year: 2024, Day: 1},
	}
	Register(2024, 1, s)

	got := Get(2024, 2)
	if got != nil {
		t.Errorf("Get(2024, 2) = %v, want nil", got)
	}
}

func TestGetYears(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	Register(2023, 1, &mockSolver{Base: solver.Base{Year: 2023, Day: 1}})
	Register(2024, 1, &mockSolver{Base: solver.Base{Year: 2024, Day: 1}})

	years := GetYears()
	if len(years) != 2 {
		t.Errorf("GetYears() returned %d years, want 2", len(years))
	}

	yearSet := make(map[int]bool)
	for _, y := range years {
		yearSet[y] = true
	}

	if !yearSet[2023] || !yearSet[2024] {
		t.Errorf("GetYears() = %v, want [2023, 2024]", years)
	}
}

func TestGetDays(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	Register(2024, 1, &mockSolver{Base: solver.Base{Year: 2024, Day: 1}})
	Register(2024, 5, &mockSolver{Base: solver.Base{Year: 2024, Day: 5}})
	Register(2024, 10, &mockSolver{Base: solver.Base{Year: 2024, Day: 10}})

	days := GetDays(2024)
	if len(days) != 3 {
		t.Errorf("GetDays(2024) returned %d days, want 3", len(days))
	}

	daySet := make(map[int]bool)
	for _, d := range days {
		daySet[d] = true
	}

	if !daySet[1] || !daySet[5] || !daySet[10] {
		t.Errorf("GetDays(2024) = %v, want [1, 5, 10]", days)
	}
}

func TestGetDaysNonExistentYear(t *testing.T) {
	// Clear registry before test
	solvers = make(map[int]map[int]solver.Solver)

	days := GetDays(2099)
	if days != nil {
		t.Errorf("GetDays(2099) = %v, want nil", days)
	}
}
