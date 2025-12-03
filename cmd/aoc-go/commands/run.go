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
		p := part.Normalize(args[2])
		if p == "" {
			log.Fatalf("Invalid part: %s (must be '1', '2', 'first', or 'second')", args[2])
		}

		// Check if we're in a project directory with main.go
		if _, err := os.Stat("main.go"); err == nil {
			runProjectSolution(year, day, p)
			return
		}

		log.Fatal("No main.go found in current directory. Please run this command from your aoc-go project root.")
	},
}

func init() {
	RunCmd.Flags().BoolP("save-stats", "s", true, "Save statistics to stats.json")
}

func runProjectSolution(year, day int, p string) {
	// Convert part to numeric for the project's main.go
	partNum := "1"
	if p == "second" {
		partNum = "2"
	}

	// Start timer
	start := time.Now()

	// Run the project's main.go
	cmd := exec.Command("go", "run", ".", strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	elapsed := time.Since(start)

	// Check stderr for compilation errors
	stderrStr := stderr.String()
	if strings.Contains(stderrStr, "build") || strings.Contains(stderrStr, "cannot find") || strings.Contains(stderrStr, "undefined") {
		fmt.Fprintf(os.Stderr, "%s", stderrStr)
		os.Exit(1)
	}

	// Check if it's a "no solution found" message
	stdoutStr := stdout.String()
	if strings.Contains(stdoutStr, "No solution found") {
		fmt.Printf("\nNo solution registered for year %d day %d part %s\n", year, day, p)
		fmt.Println("Make sure the solution is imported in main.go")
		os.Exit(1)
	}

	if err != nil {
		// Show any error output
		if stderrStr != "" {
			fmt.Fprintf(os.Stderr, "%s", stderrStr)
		}
		if stdoutStr != "" {
			fmt.Print(stdoutStr)
		}
		os.Exit(1)
	}

	// Parse the output to extract just the answer
	answer := parseAnswer(stdoutStr)

	fmt.Printf("\nAnswer: %v\n", answer)
	fmt.Printf("Time:   %s\n\n", formatExecutionTime(elapsed.Microseconds()))
}

func parseAnswer(output string) string {
	// The project's main.go outputs "Answer: <value>"
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Answer: ") {
			return strings.TrimPrefix(line, "Answer: ")
		}
	}
	return strings.TrimSpace(output)
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
