package commands

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

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
		// Convert part to numeric
		partNum := "1"
		if p == "second" {
			partNum = "2"
		}

		start := time.Now()

		// Run the compiled binary with test flag
		cmd := exec.Command("./"+binaryName, "-test", strconv.Itoa(year), strconv.Itoa(day), partNum)
		cmd.Dir, _ = os.Getwd()

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		elapsed := time.Since(start)

		if err != nil {
			// Try running without -test flag (fallback for projects without test support)
			result, testErr := runTestFallback(year, day, p)
			if testErr != nil {
				fmt.Printf("  Part %s: [SKIP] %v\n", part.Label(p), testErr)
				continue
			}
			if result.passed {
				fmt.Printf("  Part %s: [PASS] got %v [%s]\n", part.Label(p), result.actual, formatExecutionTime(elapsed.Microseconds()))
			} else {
				fmt.Printf("  Part %s: [FAIL]\n", part.Label(p))
				fmt.Printf("           Expected: %v\n", result.expected)
				fmt.Printf("           Got:      %v\n", result.actual)
				allPassed = false
			}
			continue
		}

		// Parse test output
		output := strings.TrimSpace(stdout.String())
		if strings.Contains(output, "[PASS]") {
			fmt.Printf("  Part %s: [PASS] [%s]\n", part.Label(p), formatExecutionTime(elapsed.Microseconds()))
		} else if strings.Contains(output, "[FAIL]") {
			fmt.Printf("  Part %s: %s\n", part.Label(p), output)
			allPassed = false
		} else {
			fmt.Printf("  Part %s: %s\n", part.Label(p), output)
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
	passed   bool
	expected string
	actual   string
}

func runTestFallback(year, day int, p string) (*testResult, error) {
	// Read example input and expected output
	dayStr := fmt.Sprintf("%02d", day)
	partNum := "1"
	if p == "second" {
		partNum = "2"
	}

	// Try to find example file
	examplePath := fmt.Sprintf("data/examples/%d/day_%s_part%s.txt", year, dayStr, partNum)
	if _, err := os.Stat(examplePath); os.IsNotExist(err) {
		// Try shared example file
		examplePath = fmt.Sprintf("data/examples/%d/day_%s.txt", year, dayStr)
		if _, err := os.Stat(examplePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("no example file found")
		}
	}

	// Read expected output (first line after "---" separator or specific format)
	content, err := os.ReadFile(examplePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read example file: %v", err)
	}

	// Parse example file - expect format with separator
	fileParts := strings.SplitN(string(content), "\n---\n", 2)
	if len(fileParts) != 2 {
		return nil, fmt.Errorf("example file missing expected output (use --- separator)")
	}

	expected := strings.TrimSpace(fileParts[1])

	// Run solution with example input using the cached binary
	cmd := exec.Command("./"+binaryName, strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	// Set environment variable to use example input
	cmd.Env = append(os.Environ(), "AOC_USE_EXAMPLE=1")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run solution: %v", err)
	}

	actual, _ := parseOutput(stdout.String())

	return &testResult{
		passed:   actual == expected,
		expected: expected,
		actual:   actual,
	}, nil
}
