package tester

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AnasImloul/aoc-go/pkg/registry"
)

// TestResult contains the result of running a test
type TestResult struct {
	Passed   bool
	Expected string
	Actual   any
}

// Test runs the solution for the given year/day/part against the example input
func Test(year, day int, part string) (*TestResult, error) {
	// Get the solver from registry
	solver := registry.Get(year, day)
	if solver == nil {
		return nil, fmt.Errorf("no solver registered for year %d day %d", year, day)
	}

	// Read example input
	exampleInput, err := readExampleInput(year, day)
	if err != nil {
		return nil, fmt.Errorf("could not read example input: %w", err)
	}

	// Read expected output
	expected, err := readExpectedOutput(year, day, part)
	if err != nil {
		return nil, fmt.Errorf("could not read expected output: %w", err)
	}

	// Create a temporary input file with example content
	tempDir, err := os.MkdirTemp("", "aoc-test")
	if err != nil {
		return nil, fmt.Errorf("could not create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Override the input for testing by setting an environment variable
	// that the input package can check
	os.Setenv("AOC_TEST_INPUT", exampleInput)
	defer os.Unsetenv("AOC_TEST_INPUT")

	// Run the solver
	actual := solver.Solve(part)

	// Compare results
	actualStr := fmt.Sprintf("%v", actual)
	passed := strings.TrimSpace(actualStr) == strings.TrimSpace(expected)

	return &TestResult{
		Passed:   passed,
		Expected: expected,
		Actual:   actual,
	}, nil
}

// readExampleInput reads the example input file for the given year and day
func readExampleInput(year, day int) (string, error) {
	filename := filepath.Join("data", "examples", fmt.Sprintf("%d", year), fmt.Sprintf("day_%02d.txt", day))
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// readExpectedOutput reads the expected output for the given year, day, and part
func readExpectedOutput(year, day int, part string) (string, error) {
	partNum := "1"
	if part == "second" {
		partNum = "2"
	}
	filename := filepath.Join("data", "examples", fmt.Sprintf("%d", year), fmt.Sprintf("day_%02d_part%s.txt", day, partNum))
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
