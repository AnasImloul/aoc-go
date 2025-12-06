// Package build provides utilities for building Go projects.
package build

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const (
	// BinaryName is the name of the temporary binary created during build.
	BinaryName = ".aoc-runner"
)

// BuildProject builds the current Go project and outputs a binary with the given name.
// Returns an error if the build fails.
func BuildProject(binaryName string) error {
	buildCmd := exec.Command("go", "build", "-o", binaryName, ".")
	buildCmd.Dir, _ = os.Getwd()

	var buildStderr bytes.Buffer
	buildCmd.Stderr = &buildStderr

	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("build error: %w\n%s", err, buildStderr.String())
	}
	return nil
}





