package commands

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*.tmpl
var initTemplates embed.FS

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

		// Get version from root command or build info
		version := getVersion(cmd)

		if err := initProject(projectName, moduleName, version); err != nil {
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

func initProject(projectName, moduleName, version string) error {
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

	// Template data
	templateData := struct {
		ModuleName  string
		Version     string
		ProjectName string
	}{
		ModuleName:  moduleName,
		Version:     version,
		ProjectName: projectName,
	}

	// Create files from templates
	files := []struct {
		TemplatePath string
		OutputPath   string
	}{
		{"templates/go.mod.tmpl", "go.mod"},
		{"templates/main.go.tmpl", "main.go"},
		{"templates/env.example.tmpl", ".env.example"},
		{"templates/env.example.tmpl", ".env"},
		{"templates/.gitignore.tmpl", ".gitignore"},
		{"templates/README.md.tmpl", "README.md"},
	}

	for _, file := range files {
		if err := createFileFromTemplate(initTemplates, file.TemplatePath, filepath.Join(projectName, file.OutputPath), templateData); err != nil {
			return fmt.Errorf("failed to create %s: %w", file.OutputPath, err)
		}
	}

	fmt.Printf("Created project structure for '%s'\n", projectName)
	return nil
}

// getVersion retrieves the version of aoc-go
// For released binaries: uses root command version (set by goreleaser via -X main.version)
// For dev builds: tries to get latest git tag, falls back to latest known release
func getVersion(cmd *cobra.Command) string {
	root := cmd.Root()
	
	// Try to get from root command (set by goreleaser via -X main.version)
	// This works for released binaries
	if root != nil && root.Version != "" && root.Version != "dev" {
		// Ensure it has 'v' prefix
		version := strings.TrimPrefix(root.Version, "v")
		return "v" + version
	}

	// For dev builds, try to get the latest git tag
	// This ensures we pin to a specific version for reproducible builds
	if version := getLatestGitTag(); version != "" {
		return version
	}

	// Fallback to latest known release if git is not available
	// Update this when releasing new versions
	return "v0.1.4"
}

// getLatestGitTag tries to get the latest git tag from the repository
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

// findSourceDir tries to find the aoc-go source directory
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

// createFileFromTemplate creates a file from an embedded template
func createFileFromTemplate(fs embed.FS, templatePath, outputPath string, data interface{}) error {
	templateContent, err := fs.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}


