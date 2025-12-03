package commands

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AnasImloul/aoc-go/pkg/runner"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <year> <day> <part>",
	Short: "Run a solution for a specific day and part",
	Long:  `Run a solution for a specific Advent of Code day and part (1/first or 2/second).`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
		day, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Invalid day: %v", err)
		}
		part := normalizePart(args[2])
		if part == "" {
			log.Fatalf("Invalid part: %s (must be '1', '2', 'first', or 'second')", args[2])
		}

		// Start timer
		start := time.Now()

		// Execute solution
		result := runner.Solution(year, day, part)

		// End timer
		elapsed := time.Since(start)

		fmt.Printf("\nAnswer: %v\n", result)
		fmt.Printf("Time:   %s\n\n", formatExecutionTime(elapsed.Microseconds()))
	},
}

func init() {
	RunCmd.Flags().BoolP("save-stats", "s", true, "Save statistics to stats.json")
}

func formatExecutionTime(micros int64) string {
	if micros < 1000 {
		return fmt.Sprintf("%d μs", micros)
	} else if micros < 1000000 {
		return fmt.Sprintf("%.2f ms", float64(micros)/1000.0)
	} else {
		return fmt.Sprintf("%.2f s", float64(micros)/1000000.0)
	}
}

// normalizePart converts part aliases to the canonical form
func normalizePart(part string) string {
	switch part {
	case "1", "first":
		return "first"
	case "2", "second":
		return "second"
	default:
		return ""
	}
}


