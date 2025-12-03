package commands

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AnasImloul/aoc-go/internal/part"
	"github.com/AnasImloul/aoc-go/pkg/tester"
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

		fmt.Printf("Testing Year %d Day %d\n\n", year, day)

		allPassed := true
		for _, p := range parts {
			start := time.Now()
			result, err := tester.Test(year, day, p)
			elapsed := time.Since(start)

			if err != nil {
				fmt.Printf("  Part %s: [SKIP] %v\n", part.Label(p), err)
				continue
			}

			if result.Passed {
				fmt.Printf("  Part %s: [PASS] got %v [%s]\n", part.Label(p), result.Actual, formatExecutionTime(elapsed.Microseconds()))
			} else {
				fmt.Printf("  Part %s: [FAIL]\n", part.Label(p))
				fmt.Printf("           Expected: %v\n", result.Expected)
				fmt.Printf("           Got:      %v\n", result.Actual)
				allPassed = false
			}
		}

		fmt.Println()
		if allPassed {
			fmt.Println("All tests passed!")
		} else {
			fmt.Println("Some tests failed.")
		}
	},
}
