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
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

const baseURL = "https://adventofcode.com"

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
	if testInput := os.Getenv("AOC_TEST_INPUT"); testInput != "" {
		return testInput, nil
	}

	filename := fmt.Sprintf("data/inputs/%d/day_%02d.txt", year, day)
	data, err := os.ReadFile(filename)
	if err != nil {
		input, fetchErr := fetchAndSaveInput(year, day, filename)
		if fetchErr != nil {
			return "", fetchErr
		}
		return input, nil
	}
	return strings.TrimSpace(string(data)), nil
}

// ReadLines returns a channel that streams lines from the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it streams lines from that instead (for testing).
// Errors are sent to the returned error channel.
func ReadLines(year, day int) <-chan string {
	lines, _ := TryReadLines(year, day)
	return lines
}

// TryReadLines returns a channel that streams lines from the input file for the given year and day.
// If AOC_TEST_INPUT env var is set, it streams lines from that instead (for testing).
// Returns a lines channel and an error channel. Errors during reading are sent to the error channel.
func TryReadLines(year, day int) (<-chan string, <-chan error) {
	lines := make(chan string, 100) // Buffered channel for better performance
	errChan := make(chan error, 1)

	go func() {
		defer close(lines)
		defer close(errChan)

		// Check for test input override
		if testInput := os.Getenv("AOC_TEST_INPUT"); testInput != "" {
			for _, line := range strings.Split(testInput, "\n") {
				lines <- line
			}
			return
		}

		filename := fmt.Sprintf("data/inputs/%d/day_%02d.txt", year, day)
		file, err := os.Open(filename)
		if err != nil {
			errChan <- fmt.Errorf("failed to open file %s: %w", filename, err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			errChan <- fmt.Errorf("error reading file %s: %w", filename, err)
		}
	}()

	return lines, errChan
}

// fetchAndSaveInput fetches the input from the Advent of Code website and saves it locally.
func fetchAndSaveInput(year, day int, filename string) (string, error) {
	// Load the environment variables (ignore error if .env doesn't exist)
	_ = godotenv.Load()

	// Get the session cookie from the environment variables
	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		return "", fmt.Errorf("SESSION_COOKIE is not set in environment variables; add it to your .env file")
	}

	// Construct the URL for the input
	url := fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add the session cookie for authentication
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors in the response
	if resp.StatusCode != http.StatusOK {
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

	return strings.TrimSpace(input), nil
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
