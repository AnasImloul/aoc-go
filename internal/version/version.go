// Package version provides utilities for determining the aoc-go version.
package version

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetVersion retrieves the version of aoc-go.
// For released binaries: uses root command version (set by goreleaser via -X main.version).
// For dev builds: tries to get latest git tag, then queries Go module proxy for latest published version.
// Returns an error if version cannot be determined.
func GetVersion(rootVersion string) (string, error) {
	// Try to get from root command (set by goreleaser via -X main.version)
	// This works for released binaries
	if rootVersion != "" && rootVersion != "dev" {
		// Ensure it has 'v' prefix
		version := strings.TrimPrefix(rootVersion, "v")
		return "v" + version, nil
	}

	// For dev builds, try to get the latest git tag
	// This ensures we pin to a specific version for reproducible builds
	if version := getLatestGitTag(); version != "" {
		return version, nil
	}

	// Fallback: query Go module proxy for latest published version
	if version := getLatestModuleVersion(); version != "" {
		return version, nil
	}

	// If we can't determine the version, return an error
	return "", fmt.Errorf("unable to get version from binary, git tag, or module proxy")
}

// getLatestGitTag tries to get the latest git tag from the repository.
func getLatestGitTag() string {
	// Try to find the aoc-go source directory
	// Check if we're in a git repository
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	cmd.Dir = findSourceDir()
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	tag := strings.TrimSpace(string(output))
	if tag != "" && strings.HasPrefix(tag, "v") {
		return tag
	}
	return ""
}

// getLatestModuleVersion queries the Go module proxy for the latest published version.
func getLatestModuleVersion() string {
	// Use go list to query the module proxy for latest version
	// The -f '{{.Version}}' template extracts just the version string
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/AnasImloul/aoc-go@latest")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	version := strings.TrimSpace(string(output))
	// Ensure it has 'v' prefix (should already have it, but be safe)
	if version != "" && !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	return version
}

// findSourceDir tries to find the aoc-go source directory.
func findSourceDir() string {
	// Try to find go.mod that contains aoc-go module
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Check current directory and parent directories
	for i := 0; i < 10; i++ {
		goModPath := filepath.Join(wd, "go.mod")
		if data, err := os.ReadFile(goModPath); err == nil {
			if strings.Contains(string(data), "module github.com/AnasImloul/aoc-go") {
				return wd
			}
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			break // Reached root
		}
		wd = parent
	}

	// Fallback to current directory
	return "."
}
