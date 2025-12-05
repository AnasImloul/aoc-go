package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/AnasImloul/aoc-go/internal/part"
	"github.com/spf13/cobra"
)

const binaryName = ".aoc-runner"

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

func buildProject() error {
	buildCmd := exec.Command("go", "build", "-o", binaryName, ".")
	buildCmd.Dir, _ = os.Getwd()

	var buildStderr bytes.Buffer
	buildCmd.Stderr = &buildStderr

	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("build error:\n%s", buildStderr.String())
	}
	return nil
}

type resultData struct {
	Answer     string `json:"answer"`
	TimeMicros int64  `json:"time_micros"`
}

func runProjectSolution(year, day int, p string) {
	// Convert part to numeric for the project's main.go
	partNum := "1"
	if p == "second" {
		partNum = "2"
	}

	// Build the project
	if err := buildProject(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// Clean up binary after we're done
	defer os.Remove(binaryName)

	// Create temporary result file
	resultFile, err := os.CreateTemp("", "aoc-result-*.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create result file: %v\n", err)
		os.Exit(1)
	}
	resultFilePath := resultFile.Name()
	resultFile.Close()
	defer os.Remove(resultFilePath)

	// Run the compiled binary
	cmd := exec.Command("./"+binaryName, strconv.Itoa(year), strconv.Itoa(day), partNum)
	cmd.Dir, _ = os.Getwd()

	// Pass result file path via environment variable
	cmd.Env = append(os.Environ(), "AOC_RESULT_FILE="+resultFilePath)

	// Write stdout/stderr directly to terminal for user logs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	// Check if execution failed
	if err != nil {
		// Check if it's a "no solution found" message by reading stderr
		// (we can't capture it separately now, but main.go should exit with code 1)
		os.Exit(1)
	}

	// Read result from file
	resultData, err := readResultFile(resultFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read result file: %v\n", err)
		os.Exit(1)
	}

	// Display answer and timing
	if resultData.Answer != "" {
		fmt.Printf("\nAnswer: %s\n", resultData.Answer)
		fmt.Printf("Time:   %s\n\n", formatExecutionTime(resultData.TimeMicros))
	}
}

func readResultFile(filePath string) (*resultData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result resultData
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse result JSON: %w", err)
	}

	return &result, nil
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
