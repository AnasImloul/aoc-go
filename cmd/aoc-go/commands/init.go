package commands

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/AnasImloul/aoc-go/internal/constants"
	"github.com/AnasImloul/aoc-go/internal/version"
	"github.com/spf13/cobra"
)

//go:embed templates/*.tmpl
var initTemplates embed.FS

var InitCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new Advent of Code project",
	Long:  `Create a new Advent of Code project with the necessary structure and configuration.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		moduleName, _ := cmd.Flags().GetString("module")

		if moduleName == "" {
			moduleName = projectName
		}

		// Get version from root command or build info
		rootVersion := ""
		if root := cmd.Root(); root != nil {
			rootVersion = root.Version
		}
		version, err := version.GetVersion(rootVersion)
		if err != nil {
			return fmt.Errorf("failed to determine version: %w", err)
		}

		if err := initProject(projectName, moduleName, version); err != nil {
			return fmt.Errorf("error initializing project: %w", err)
		}

		fmt.Printf("\n✓ Project '%s' created successfully!\n\n", projectName)
		fmt.Println("Next steps:")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Println("  # Add your SESSION_COOKIE to .env")
		fmt.Println("  aoc-go generate 2024 1")
		fmt.Println("  aoc-go run 2024 1 1")
		fmt.Println()
		return nil
	},
}

func init() {
	InitCmd.Flags().StringP("module", "m", "", "Go module name (defaults to project name)")
}

func initProject(projectName, moduleName, version string) error {
	if err := createProjectDirectories(projectName); err != nil {
		return err
	}

	templateData := newTemplateData(moduleName, version, projectName)
	if err := createProjectFiles(projectName, templateData); err != nil {
		return err
	}

	if err := runGoModTidy(projectName); err != nil {
		return err
	}

	fmt.Printf("Created project structure for '%s'\n", projectName)
	return nil
}

func createProjectDirectories(projectName string) error {
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	dirs := []string{
		constants.SolutionsDir,
		filepath.Join(constants.DataDir, constants.InputsDir),
		filepath.Join(constants.DataDir, constants.ExamplesDir),
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	return nil
}

func newTemplateData(moduleName, version, projectName string) struct {
	ModuleName  string
	Version     string
	ProjectName string
} {
	return struct {
		ModuleName  string
		Version     string
		ProjectName string
	}{
		ModuleName:  moduleName,
		Version:     version,
		ProjectName: projectName,
	}
}

func createProjectFiles(projectName string, templateData interface{}) error {
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
		outputPath := filepath.Join(projectName, file.OutputPath)
		if err := createFileFromTemplate(initTemplates, file.TemplatePath, outputPath, templateData); err != nil {
			return fmt.Errorf("failed to create %s: %w", file.OutputPath, err)
		}
	}
	return nil
}

func runGoModTidy(projectName string) error {
	fmt.Printf("Downloading dependencies...\n")
	modTidyCmd := exec.Command("go", "mod", "tidy")
	modTidyCmd.Dir = projectName
	if output, err := modTidyCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %w\noutput: %s", err, string(output))
	}
	return nil
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
