package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new Advent of Code project",
	Long:  `Create a new Advent of Code project with the necessary structure and configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		moduleName, _ := cmd.Flags().GetString("module")

		if moduleName == "" {
			moduleName = projectName
		}

		if err := initProject(projectName, moduleName); err != nil {
			log.Fatalf("Error initializing project: %v", err)
		}

		fmt.Printf("\n✓ Project '%s' created successfully!\n\n", projectName)
		fmt.Println("Next steps:")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Println("  # Add your SESSION_COOKIE to .env")
		fmt.Println("  aoc-go generate 2024 1")
		fmt.Println("  aoc-go run 2024 1 1")
		fmt.Println()
	},
}

func init() {
	InitCmd.Flags().StringP("module", "m", "", "Go module name (defaults to project name)")
}

func initProject(projectName, moduleName string) error {
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create subdirectories
	dirs := []string{
		"solutions",
		"data/inputs",
		"data/examples",
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}

	// Create go.mod
	goMod := fmt.Sprintf(`module %s

go 1.22

require github.com/AnasImloul/aoc-go v0.1.0
`, moduleName)

	if err := os.WriteFile(filepath.Join(projectName, "go.mod"), []byte(goMod), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	// Create main.go
	mainGo := fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AnasImloul/aoc-go/pkg/runner"

	// Import your solutions here - they register themselves via init()
	// Solutions are automatically added when you run: aoc-go generate <year> <day>
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <year> <day> <part>")
		fmt.Println("Example: go run main.go 2024 1 1")
		os.Exit(1)
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid year: %%v\n", err)
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid day: %%v\n", err)
		os.Exit(1)
	}

	part := normalizePart(os.Args[3])
	if part == "" {
		fmt.Printf("Invalid part: %%s (must be '1', '2', 'first', or 'second')\n", os.Args[3])
		os.Exit(1)
	}

	result := runner.Solution(year, day, part)
	if result == nil {
		fmt.Printf("No solution found for year %%d day %%d\n", year, day)
		os.Exit(1)
	}

	fmt.Printf("Answer: %%v\n", result)
}

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
`)

	if err := os.WriteFile(filepath.Join(projectName, "main.go"), []byte(mainGo), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	// Create .env.example
	envExample := `# Advent of Code session cookie
# Get this from your browser's developer tools after logging into adventofcode.com
# Look for the 'session' cookie value
SESSION_COOKIE=your_session_cookie_here
`
	if err := os.WriteFile(filepath.Join(projectName, ".env.example"), []byte(envExample), 0644); err != nil {
		return fmt.Errorf("failed to create .env.example: %w", err)
	}

	// Create .env (copy of .env.example)
	if err := os.WriteFile(filepath.Join(projectName, ".env"), []byte(envExample), 0644); err != nil {
		return fmt.Errorf("failed to create .env: %w", err)
	}

	// Create .gitignore
	gitignore := `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# Environment variables
.env

# IDE
.idea/
.vscode/

# OS files
.DS_Store
Thumbs.db

# Input files (personal puzzle inputs)
data/inputs/
`
	if err := os.WriteFile(filepath.Join(projectName, ".gitignore"), []byte(gitignore), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Create README.md
	readme := fmt.Sprintf(`# %s

My Advent of Code solutions in Go, powered by [aoc-go](https://github.com/AnasImloul/aoc-go).

## Setup

1. Install the aoc-go CLI:
   `+"```bash"+`
   brew install AnasImloul/tap/aoc-go
   `+"```"+`

2. Get your session cookie from [adventofcode.com](https://adventofcode.com):
   - Log in to Advent of Code
   - Open browser developer tools (F12)
   - Go to Application/Storage > Cookies
   - Copy the value of the 'session' cookie

3. Add your session cookie to `+"`"+`.env`+"`"+`:
   `+"```bash"+`
   SESSION_COOKIE=your_session_cookie_here
   `+"```"+`

## Usage

### Generate a new day
`+"```bash"+`
aoc-go generate <year> <day>
# Example: aoc-go generate 2024 1
`+"```"+`

### Run a solution
`+"```bash"+`
go run main.go <year> <day> <part>
# Example: go run main.go 2024 1 1
`+"```"+`

Or use the CLI directly:
`+"```bash"+`
aoc-go run <year> <day> <part>
`+"```"+`

### Test against examples
`+"```bash"+`
aoc-go test <year> <day>
`+"```"+`

## Project Structure

`+"```"+`
%s/
├── main.go           # Entry point
├── solutions/        # Your solutions
│   └── 2024/
│       └── day01/
│           ├── solution.go
│           ├── part1.go
│           └── part2.go
├── data/
│   ├── inputs/      # Puzzle inputs (auto-downloaded)
│   └── examples/    # Example inputs for testing
├── .env             # Session cookie (gitignored)
└── go.mod
`+"```"+`
`, projectName, projectName)

	if err := os.WriteFile(filepath.Join(projectName, "README.md"), []byte(readme), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}

	fmt.Printf("Created project structure for '%s'\n", projectName)
	return nil
}


