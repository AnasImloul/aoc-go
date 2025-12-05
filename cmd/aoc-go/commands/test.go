package commands

import (
	_ "encoding/json" // Used by readResultFile in run.go
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AnasImloul/aoc-go/internal/part"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test <year> <day> [part]",
	Short: "Test a solution against example input",
	Long:  `Test a solution for a specific Advent of Code day against the example input and expected output.`,
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
		day, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Invalid day: %v", err)
		}

		// Determine which parts to test
		parts := []string{"first", "second"}
		if len(args) == 3 {
			p := part.Normalize(args[2])
			if p == "" {
				log.Fatalf("Invalid part: %s (must be '1', '2', 'first', or 'second')", args[2])
			}
			parts = []string{p}
		}

		// Check if we're in a project directory with main.go
		if _, err := os.Stat("main.go"); err == nil {
			runProjectTests(year, day, parts)
			return
		}

		log.Fatal("No main.go found in current directory. Please run this command from your aoc-go project root.")
	},
}

func runProjectTests(year, day int, parts []string) {
	fmt.Printf("Testing Year %d Day %d\n\n", year, day)

	// Build the project
	if err := buildProject(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// Clean up binary after we're done
	defer os.Remove(binaryName)

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
			fmt.Printf("  Part %s: [PASS] got %v [%s]\n", part.Label(p), result.actual, formatExecutionTime(result.timeMicros))
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
	dayFolderPath := fmt.Sprintf("data/examples/%d/%s", year, dayFolderName)
	inputFile := filepath.Join(dayFolderPath, "input.txt")
	partFile := filepath.Join(dayFolderPath, fmt.Sprintf("part%s.txt", partNum))

	inputContent, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("example input file not found: %s", inputFile)
	}
	exampleInput = strings.TrimSpace(string(inputContent))

	partContent, err := os.ReadFile(partFile)
	if err != nil {
		return nil, fmt.Errorf("expected output file not found: %s", partFile)
	}
	expected = strings.TrimSpace(string(partContent))

	if exampleInput == "" {
		return nil, fmt.Errorf("no example input found")
	}
	if expected == "" {
		return nil, fmt.Errorf("no expected output found")
	}

	// Create temporary result file
	resultFile, err := os.CreateTemp("", "aoc-result-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create result file: %v", err)
	}
	resultFilePath := resultFile.Name()
	resultFile.Close()
	defer os.Remove(resultFilePath)

	// Run solution with example input using the cached binary
	cmd := exec.Command("./"+binaryName, strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	// Set environment variables
	cmd.Env = append(os.Environ(), "AOC_TEST_INPUT="+exampleInput, "AOC_RESULT_FILE="+resultFilePath)

	// Write stdout/stderr directly to terminal for user logs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run solution: %v", err)
	}

	// Read result from file
	resultData, err := readResultFile(resultFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read result file: %v", err)
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
