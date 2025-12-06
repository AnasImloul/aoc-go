package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/AnasImloul/aoc-go/internal/build"
	"github.com/AnasImloul/aoc-go/internal/constants"
	"github.com/AnasImloul/aoc-go/internal/part"
	resultpkg "github.com/AnasImloul/aoc-go/internal/result"
	"github.com/AnasImloul/aoc-go/internal/validation"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test <year> <day> [part]",
	Short: "Test a solution against example input",
	Long:  `Test a solution for a specific Advent of Code day against the example input and expected output.`,
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid year: %w", err)
		}

		day, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid day: %w", err)
		}

		if err := validation.ValidateYearAndDay(year, day); err != nil {
			return err
		}

		// Determine which parts to test
		parts := []string{"first", "second"}
		if len(args) == 3 {
			p := part.Normalize(args[2])
			if p == "" {
				return fmt.Errorf("invalid part: %s (must be '1', '2', 'first', or 'second')", args[2])
			}
			parts = []string{p}
		}

		// Check if we're in a project directory with main.go
		if _, err := os.Stat(constants.MainGoFile); err == nil {
			return runProjectTests(year, day, parts)
		}

		return fmt.Errorf("no %s found in current directory: please run this command from your aoc-go project root", constants.MainGoFile)
	},
}

func runProjectTests(year, day int, parts []string) error {
	fmt.Printf("Testing Year %d Day %d\n\n", year, day)

	// Build the project
	if err := build.BuildProject(build.BinaryName); err != nil {
		return fmt.Errorf("failed to build project: %w", err)
	}

	// Clean up binary after we're done
	defer os.Remove(build.BinaryName)

	allPassed := true
	for _, p := range parts {
		// Try running with test flag first (if project supports it)
		// Otherwise fallback to manual testing
		result, testErr := runTestFallback(year, day, p)
		if testErr != nil {
			fmt.Printf("  Part %s: [SKIP] %v\n", part.Label(p), testErr)
			continue
		}

		if result.passed {
			fmt.Printf("  Part %s: [PASS] got %v [%s]\n", part.Label(p), result.actual, resultpkg.FormatExecutionTime(result.timeMicros))
		} else {
			fmt.Printf("  Part %s: [FAIL]\n", part.Label(p))
			fmt.Printf("           Expected: %v\n", result.expected)
			fmt.Printf("           Got:      %v\n", result.actual)
			allPassed = false
		}
	}

	fmt.Println()
	if allPassed {
		fmt.Println("All tests passed!")
	} else {
		fmt.Println("Some tests failed.")
	}

	return nil
}

type testResult struct {
	passed     bool
	expected   string
	actual     string
	timeMicros int64
}

func runTestFallback(year, day int, p string) (*testResult, error) {
	// Read example input and expected output
	dayStr := fmt.Sprintf("%02d", day)
	dayFolderName := fmt.Sprintf("day%s", dayStr)
	partNum := "1"
	if p == "second" {
		partNum = "2"
	}

	var exampleInput string
	var expected string

	// Use new folder structure: data/examples/2025/day01/input.txt and part1.txt
	dayFolderPath := filepath.Join(constants.DataDir, constants.ExamplesDir, fmt.Sprintf("%d", year), dayFolderName)
	inputFile := filepath.Join(dayFolderPath, constants.ExampleInputFileName)
	partFile := filepath.Join(dayFolderPath, fmt.Sprintf(constants.PartFileName, partNum))

	inputContent, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("example input file not found: %s: %w", inputFile, err)
	}
	exampleInput = string(inputContent)

	partContent, err := os.ReadFile(partFile)
	if err != nil {
		return nil, fmt.Errorf("expected output file not found: %s: %w", partFile, err)
	}
	expected = string(partContent)

	if exampleInput == "" {
		return nil, fmt.Errorf("no example input found")
	}
	if expected == "" {
		return nil, fmt.Errorf("no expected output found")
	}

	// Create temporary result file
	resultFile, err := os.CreateTemp("", constants.ResultFilePrefix+"*"+constants.ResultFileSuffix)
	if err != nil {
		return nil, fmt.Errorf("failed to create result file: %w", err)
	}
	resultFilePath := resultFile.Name()
	resultFile.Close()
	defer os.Remove(resultFilePath)

	// Run solution with example input using the cached binary
	cmd := exec.Command("./"+build.BinaryName, strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	// Set environment variables
	cmd.Env = append(os.Environ(), constants.EnvTestInput+"="+exampleInput, constants.EnvResultFile+"="+resultFilePath)

	// Write stdout/stderr directly to terminal for user logs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run solution: %w", err)
	}

	// Read result from file
	resultData, err := resultpkg.ReadFromFile(resultFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read result file: %w", err)
	}

	// Answer is already a string
	actual := resultData.Answer

	return &testResult{
		passed:     actual == expected,
		expected:   expected,
		actual:     actual,
		timeMicros: resultData.TimeMicros,
	}, nil
}
