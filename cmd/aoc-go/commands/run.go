package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/AnasImloul/aoc-go/internal/build"
	"github.com/AnasImloul/aoc-go/internal/constants"
	"github.com/AnasImloul/aoc-go/internal/part"
	"github.com/AnasImloul/aoc-go/internal/result"
	"github.com/AnasImloul/aoc-go/internal/validation"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <year> <day> <part>",
	Short: "Run a solution for a specific day and part",
	Long:  `Run a solution for a specific Advent of Code day and part (1/first or 2/second).`,
	Args:  cobra.ExactArgs(3),
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

		p := part.Normalize(args[2])
		if p == "" {
			return fmt.Errorf("invalid part: %s (must be '1', '2', 'first', or 'second')", args[2])
		}

		// Check if we're in a project directory with main.go
		if _, err := os.Stat(constants.MainGoFile); err == nil {
			return runProjectSolution(year, day, p)
		}

		return fmt.Errorf("no %s found in current directory: please run this command from your aoc-go project root", constants.MainGoFile)
	},
}

func init() {
	RunCmd.Flags().BoolP("save-stats", "s", true, "Save statistics to stats.json")
}

type resultData struct {
	Answer     string `json:"answer"`
	TimeMicros int64  `json:"time_micros"`
}

func runProjectSolution(year, day int, p string) error {
	// Convert part to numeric for the project's main.go
	partNum := "1"
	if p == "second" {
		partNum = "2"
	}

	// Build the project
	if err := build.BuildProject(build.BinaryName); err != nil {
		return fmt.Errorf("failed to build project: %w", err)
	}

	// Clean up binary after we're done
	defer os.Remove(build.BinaryName)

	// Create temporary result file
	resultFile, err := os.CreateTemp("", constants.ResultFilePrefix+"*"+constants.ResultFileSuffix)
	if err != nil {
		return fmt.Errorf("failed to create result file: %w", err)
	}
	resultFilePath := resultFile.Name()
	resultFile.Close()
	defer os.Remove(resultFilePath)

	// Run the compiled binary
	cmd := exec.Command("./"+build.BinaryName, strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	// Pass result file path via environment variable
	cmd.Env = append(os.Environ(), constants.EnvResultFile+"="+resultFilePath)

	// Write stdout/stderr directly to terminal for user logs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("solution execution failed: %w", err)
	}

	// Read result from file
	resultData, err := result.ReadFromFile(resultFilePath)
	if err != nil {
		return fmt.Errorf("failed to read result file: %w", err)
	}

	// Display answer and timing
	if resultData.Answer != "" {
		fmt.Printf("\nAnswer: %s\n", resultData.Answer)
		fmt.Printf("Time:   %s\n\n", result.FormatExecutionTime(resultData.TimeMicros))
	}

	return nil
}
