# aoc-go

A CLI tool for managing and running Advent of Code solutions in Go.

## Features

- **Project scaffolding**: Quickly initialize a new AoC project with the correct structure
- **Solution generation**: Generate boilerplate code for each day with various templates
- **Input fetching**: Automatically download puzzle inputs from adventofcode.com
- **Testing**: Run solutions against example inputs
- **Cross-platform**: Works on macOS, Linux, and Windows

## Installation

### Homebrew (macOS/Linux)

```bash
brew install AnasImloul/tap/aoc-go
```

### Go Install

```bash
go install github.com/AnasImloul/aoc-go/cmd/aoc-go@latest
```

### From Source

```bash
git clone https://github.com/AnasImloul/aoc-go.git
cd aoc-go
go build -o aoc-go ./cmd/aoc-go
```

## Quick Start

### 1. Create a New Project

```bash
aoc-go init my-aoc-solutions
cd my-aoc-solutions
```

### 2. Configure Session Cookie

Get your session cookie from [adventofcode.com](https://adventofcode.com):
1. Log in to Advent of Code
2. Open browser developer tools (F12)
3. Go to Application/Storage > Cookies
4. Copy the value of the 'session' cookie

Add it to your `.env` file:
```bash
SESSION_COOKIE=your_session_cookie_here
```

### 3. Generate a Day

```bash
aoc-go generate 2024 1
```

This creates:
```
solutions/2024/day01/
├── solution.go
├── part1.go
└── part2.go
```

### 4. Solve the Puzzle

Edit `solutions/2024/day01/part1.go` to implement your solution:

```go
func (d *Day) firstPart() any {
    // Your solution here
    return result
}
```

### 5. Run Your Solution

```bash
# Using your project's main.go
go run main.go 2024 1 1

# Or using the CLI (after importing solutions in main.go)
aoc-go run 2024 1 1
```

## Commands

### `aoc-go init <project-name>`

Initialize a new Advent of Code project.

```bash
aoc-go init my-solutions
aoc-go init my-solutions --module github.com/username/my-solutions
```

Options:
- `-m, --module`: Go module name (defaults to project name)

### `aoc-go generate <year> <day>`

Generate template files for a new day.

```bash
aoc-go generate 2024 1
aoc-go generate 2024 1 --pattern grid
aoc-go generate 2024 1 --pattern graph --example=false
```

Options:
- `-p, --pattern`: Template pattern (`grid`, `graph`, `parser`, `simulation`)
- `-e, --example`: Generate example files (default: true)

### `aoc-go run <year> <day> <part>`

Run a solution for a specific day and part.

```bash
aoc-go run 2024 1 1
aoc-go run 2024 1 2
```

### `aoc-go test <year> <day> [part]`

Test a solution against example input.

```bash
aoc-go test 2024 1      # Test both parts
aoc-go test 2024 1 1    # Test part 1 only
```

## Template Patterns

The `generate` command supports several patterns for common puzzle types:

### Default
Basic template with minimal boilerplate.

### Grid (`--pattern grid`)
For puzzles involving 2D grids:
- Pre-parsed grid as `[]string`
- `inBounds()` and `getCell()` helpers

### Graph (`--pattern graph`)
For graph traversal puzzles:
- Adjacency list representation
- BFS and DFS implementations

### Parser (`--pattern parser`)
For puzzles requiring complex input parsing:
- Tokenization helpers
- Structured data parsing

### Simulation (`--pattern simulation`)
For step-based simulation puzzles:
- State management
- `step()` and `simulate()` methods

## Project Structure

A generated project looks like this:

```
my-aoc-solutions/
├── main.go              # Entry point
├── solutions/           # Your solutions
│   └── 2024/
│       └── day01/
│           ├── solution.go
│           ├── part1.go
│           └── part2.go
├── data/
│   ├── inputs/         # Puzzle inputs (auto-downloaded)
│   └── examples/       # Example inputs for testing
├── .env                # Session cookie (gitignored)
├── .gitignore
├── go.mod
└── README.md
```

## Using as a Library

You can use aoc-go packages directly in your code:

```go
import (
    "github.com/AnasImloul/aoc-go/pkg/solver"
    "github.com/AnasImloul/aoc-go/pkg/registry"
    "github.com/AnasImloul/aoc-go/pkg/utils"
)
```

### Packages

- `pkg/solver`: Solver interface and Base struct for solutions
- `pkg/registry`: Global solution registry
- `pkg/input`: Input fetching and reading
- `pkg/utils`: Utility functions (parsing, grids, functional helpers)
- `pkg/generator`: Solution file generation
- `pkg/runner`: Solution execution
- `pkg/tester`: Testing utilities

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.
