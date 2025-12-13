// Package input provides utilities for reading Advent of Code puzzle inputs.
//
// Inputs are cached locally in the data/inputs directory. If a cached input
// does not exist, it is automatically fetched from adventofcode.com using
// the SESSION_COOKIE environment variable for authentication.
//
// # Thread Safety
//
// Read operations are thread-safe for reading cached files. However,
// fetching and caching new inputs involves file I/O that may race with
// concurrent reads of the same file. For concurrent access to inputs
// that may not be cached, synchronize access externally.
//
// The TryReadLines function returns channels that are safe to consume
// from a single goroutine.
package input

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AnasImloul/aoc-go/internal/constants"
	"github.com/joho/godotenv"
)

const baseURL = constants.AdventOfCodeBaseURL

// trimTrailingEmptyLines removes trailing lines that contain only whitespace.
// Lines with actual data but trailing spaces are preserved.
func trimTrailingEmptyLines(lines []string) []string {
	// Find the last line that contains non-whitespace characters
	lastNonEmpty := -1
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			lastNonEmpty = i
			break
		}
	}
	// If all lines are empty, return empty slice
	if lastNonEmpty == -1 {
		return []string{}
	}
	// Return lines up to and including the last non-empty line
	return lines[:lastNonEmpty+1]
}

// Read reads the content of the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it returns that instead (for testing).
// If the file does not exist, it fetches the input from the Advent of Code website and saves it.
// Panics on error - use TryRead for error handling.
func Read(year, day int) string {
	input, err := TryRead(year, day)
	if err != nil {
		panic(fmt.Sprintf("failed to read input for year %d day %d: %v", year, day, err))
	}
	return input
}

// TryRead reads the content of the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it returns that instead (for testing).
// If the file does not exist, it fetches the input from the Advent of Code website and saves it.
// Returns an error if the input cannot be read or fetched.
func TryRead(year, day int) (string, error) {
	// Check for test input override
	// Use LookupEnv to get the exact value, including trailing spaces
	if testInput, ok := os.LookupEnv(constants.EnvTestInput); ok && testInput != "" {
		// Trim trailing empty lines from test input for consistency
		lines := strings.Split(testInput, "\n")
		lines = trimTrailingEmptyLines(lines)
		return strings.Join(lines, "\n"), nil
	}

	filename := filepath.Join(constants.DataDir, constants.InputsDir, fmt.Sprintf("%d", year), fmt.Sprintf(constants.InputFileName, day))
	data, err := os.ReadFile(filename)
	if err != nil {
		input, fetchErr := fetchAndSaveInput(year, day, filename)
		if fetchErr != nil {
			return "", fetchErr
		}
		// Trim trailing empty lines from fetched input
		lines := strings.Split(input, "\n")
		lines = trimTrailingEmptyLines(lines)
		return strings.Join(lines, "\n"), nil
	}
	// Trim trailing empty lines from file content
	lines := strings.Split(string(data), "\n")
	lines = trimTrailingEmptyLines(lines)
	return strings.Join(lines, "\n"), nil
}

// ReadLines returns a channel that streams lines from the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it streams lines from that instead (for testing).
// If the file does not exist, it fetches the input from the Advent of Code website and saves it.
// Errors are sent to the returned error channel.
func ReadLines(year, day int) <-chan string {
	lines, _ := TryReadLines(year, day)
	return lines
}

// TryReadLines returns a channel that streams lines from the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it streams lines from that instead (for testing).
// If the file does not exist, it fetches the input from the Advent of Code website and saves it.
// Returns a lines channel and an error channel. Errors during reading are sent to the error channel.
// This function is a wrapper around TryRead that splits the input into lines and streams them.
func TryReadLines(year, day int) (<-chan string, <-chan error) {
	lines := make(chan string, 100) // Buffered channel for better performance
	errChan := make(chan error, 1)

	go func() {
		defer close(lines)
		defer close(errChan)

		// Reuse TryRead to get the input (handles test input, file reading, fetching, and trimming)
		input, err := TryRead(year, day)
		if err != nil {
			errChan <- err
			return
		}

		// Split into lines and stream them
		// Note: TryRead already handles trimming trailing empty lines, so we just split
		inputLines := strings.Split(input, "\n")
		for _, line := range inputLines {
			lines <- line
		}
	}()

	return lines, errChan
}

// fetchAndSaveInput fetches the input from the Advent of Code website and saves it locally.
func fetchAndSaveInput(year, day int, filename string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return fetchAndSaveInputWithContext(ctx, year, day, filename)
}

// fetchAndSaveInputWithContext fetches the input from the Advent of Code website and saves it locally.
// The context can be used to cancel the request or set a timeout.
func fetchAndSaveInputWithContext(ctx context.Context, year, day int, filename string) (string, error) {
	// Load the environment variables (ignore error if .env doesn't exist)
	_ = godotenv.Load()

	// Get the session cookie from the environment variables
	sessionCookie := os.Getenv(constants.SessionCookieEnv)
	if sessionCookie == "" {
		return "", fmt.Errorf("%s is not set in environment variables; add it to your .env file", constants.SessionCookieEnv)
	}

	// Construct the URL for the input
	url := fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day)

	// Create the HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add the session cookie for authentication
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	// Send the request with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("request timeout: %w", err)
		}
		if ctx.Err() == context.Canceled {
			return "", fmt.Errorf("request canceled: %w", err)
		}
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors in the response
	if resp.StatusCode != constants.HTTPStatusOK {
		return "", fmt.Errorf("failed to fetch input: HTTP %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	input := string(body)

	// Save the input to the local file
	if err := saveInputToFile(input, filename); err != nil {
		return "", fmt.Errorf("failed to save input: %w", err)
	}

	return input, nil
}

// saveInputToFile saves the input to the specified file.
func saveInputToFile(input, filename string) error {
	// Ensure the directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write the input to the file
	if err := os.WriteFile(filename, []byte(input), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}
