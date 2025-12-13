// Package constants provides shared constants used across the aoc-go CLI tool.
package constants

const (
	// BinaryName is the name of the temporary binary created during build.
	BinaryName = ".aoc-runner"

	// MainGoFile is the expected main.go file name in project directories.
	MainGoFile = "main.go"

	// ResultFilePrefix is the prefix for temporary result files.
	ResultFilePrefix = "aoc-result-"

	// ResultFileSuffix is the suffix for temporary result files.
	ResultFileSuffix = ".json"

	// EnvResultFile is the environment variable name for the result file path.
	EnvResultFile = "AOC_RESULT_FILE"

	// EnvTestInput is the environment variable name for test input override.
	EnvTestInput = "AOC_TEST_INPUT"

	// SessionCookieEnv is the environment variable name for the session cookie.
	SessionCookieEnv = "SESSION_COOKIE"
)

// File paths
const (
	// SolutionsDir is the directory name for solutions.
	SolutionsDir = "solutions"

	// DataDir is the directory name for data.
	DataDir = "data"

	// InputsDir is the directory name for inputs.
	InputsDir = "inputs"

	// ExamplesDir is the directory name for examples.
	ExamplesDir = "examples"

	// InputFileName is the template for input file names.
	InputFileName = "day_%02d.txt"

	// ExampleInputFileName is the name of example input files.
	ExampleInputFileName = "input.txt"

	// PartFileName is the template for part output file names.
	PartFileName = "part%s.txt"

	// DayFolderName is the template for day folder names.
	DayFolderName = "day%02s"
)

// HTTP constants
const (
	// AdventOfCodeBaseURL is the base URL for Advent of Code.
	AdventOfCodeBaseURL = "https://adventofcode.com"

	// HTTPStatusOK is the HTTP status code for successful requests.
	HTTPStatusOK = 200
)
