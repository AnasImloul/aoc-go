package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/AnasImloul/aoc-go/pkg/input"
)

// Templates holds embedded template files
//
//go:embed templates/*.txt
var Templates embed.FS

// GenerateFiles generates the solution files for a given year and day.
func GenerateFiles(year int, day int, pattern string, withExample bool) error {
	// Format year and day
	yearStr := strconv.Itoa(year)
	dayStr := fmt.Sprintf("%02d", day)

	// Define paths
	basePath := filepath.Join(".", "solutions", yearStr, fmt.Sprintf("day%s", dayStr))

	// Create the directory structure (idempotent)
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Determine the solution template to use based on pattern
	solutionTemplate := "solution.txt"
	if pattern != "" {
		patternTemplate := fmt.Sprintf("base_%s.txt", pattern)
		if _, err := Templates.ReadFile("templates/" + patternTemplate); err == nil {
			solutionTemplate = patternTemplate
		} else {
			fmt.Printf("Pattern template %s not found, using default solution.txt\n", patternTemplate)
		}
	}

	// List of template files and corresponding output files
	files := []struct {
		TemplateFile string
		OutputFile   string
	}{
		{solutionTemplate, "solution.go"},
		{"part1.txt", "part1.go"},
		{"part2.txt", "part2.go"},
	}

	// Get module name from go.mod
	moduleName, err := getModuleName()
	if err != nil {
		return fmt.Errorf("error reading module name: %w", err)
	}

	// Process each template
	for _, file := range files {
		outputFilePath := filepath.Join(basePath, file.OutputFile)

		// Skip file creation if it already exists
		if _, err := os.Stat(outputFilePath); err == nil {
			fmt.Printf("File %s already exists, skipping...\n", outputFilePath)
			continue
		}

		templateContent, err := Templates.ReadFile("templates/" + file.TemplateFile)
		if err != nil {
			return fmt.Errorf("error reading template file %s: %w", file.TemplateFile, err)
		}

		// Replace placeholders in the template
		customizedContent := strings.ReplaceAll(string(templateContent), "MODULE_NAME", moduleName)
		customizedContent = strings.ReplaceAll(customizedContent, "YEAR", yearStr)
		customizedContent = strings.ReplaceAll(customizedContent, "day_XX", fmt.Sprintf("day%s", dayStr))
		customizedContent = strings.ReplaceAll(customizedContent, "dayXX", fmt.Sprintf("day%s", dayStr))
		customizedContent = strings.ReplaceAll(customizedContent, "Day:  XX", fmt.Sprintf("Day:  %d", day))
		customizedContent = strings.ReplaceAll(customizedContent, "XX", strconv.Itoa(day))

		// Write the customized content to the output file
		if err := os.WriteFile(outputFilePath, []byte(customizedContent), 0644); err != nil {
			return fmt.Errorf("error creating file %s: %w", outputFilePath, err)
		}
	}

	// Update the main.go file to add the blank import for the new solution
	if err := updateMainFile(moduleName, year, day); err != nil {
		return fmt.Errorf("error updating main file: %w", err)
	}

	// Create example files if requested
	if withExample {
		if err := createExampleFiles(year, day); err != nil {
			fmt.Printf("Warning: could not create example files: %v\n", err)
		}
	}

	// Fetch input file
	input.Read(year, day)

	fmt.Printf("Folder and files created successfully at %s\n", basePath)
	return nil
}

func getModuleName() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", fmt.Errorf("module name not found in go.mod")
}

func createExampleFiles(year, day int) error {
	examplesPath := filepath.Join(".", "data", "examples", strconv.Itoa(year))

	// Create the directory structure
	if err := os.MkdirAll(examplesPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating examples directory: %w", err)
	}

	dayStr := fmt.Sprintf("%02d", day)

	// Create example input file
	inputFile := filepath.Join(examplesPath, fmt.Sprintf("day_%s.txt", dayStr))
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		content := "# Paste example input here\n"
		if err := os.WriteFile(inputFile, []byte(content), 0644); err != nil {
			return fmt.Errorf("error creating example input file: %w", err)
		}
		fmt.Printf("Created example input file: %s\n", inputFile)
	}

	// Create expected output files for part 1 and part 2
	for _, partNum := range []string{"1", "2"} {
		outputFile := filepath.Join(examplesPath, fmt.Sprintf("day_%s_part%s.txt", dayStr, partNum))
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			content := "# Expected output for part " + partNum + "\n"
			if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
				return fmt.Errorf("error creating expected output file: %w", err)
			}
			fmt.Printf("Created expected output file: %s\n", outputFile)
		}
	}

	return nil
}

func updateMainFile(moduleName string, year, day int) error {
	dayStr := fmt.Sprintf("%02d", day)
	mainFilePath := "main.go"

	// Check if main.go file exists
	if _, err := os.Stat(mainFilePath); err != nil {
		return fmt.Errorf("main.go file does not exist: %w", err)
	}

	// Read the existing main.go file
	content, err := os.ReadFile(mainFilePath)
	if err != nil {
		return fmt.Errorf("error reading main file: %w", err)
	}

	contentStr := string(content)

	// Prepare the blank import statement for the new solution
	importStatement := fmt.Sprintf("_ \"%s/solutions/%d/day%s\"", moduleName, year, dayStr)

	// Check if the import statement already exists
	importPattern := regexp.MustCompile(regexp.QuoteMeta(importStatement))
	if importPattern.MatchString(contentStr) {
		fmt.Printf("Import for %d/day%s already exists, skipping...\n", year, dayStr)
		return nil
	}

	// Find the import block - look for "import (" and find the matching ")"
	importStart := strings.Index(contentStr, "import (")
	if importStart == -1 {
		return fmt.Errorf("could not find import block in main.go")
	}

	// Find the closing parenthesis of the import block
	// Start searching from after "import ("
	searchStart := importStart + len("import (")
	depth := 1
	importBlockEnd := -1
	for i := searchStart; i < len(contentStr); i++ {
		if contentStr[i] == '(' {
			depth++
		} else if contentStr[i] == ')' {
			depth--
			if depth == 0 {
				importBlockEnd = i
				break
			}
		}
	}

	if importBlockEnd == -1 {
		return fmt.Errorf("could not find end of import block in main.go")
	}

	// Insert the new import before the closing parenthesis
	contentStr = contentStr[:importBlockEnd] + "\t" + importStatement + "\n" + contentStr[importBlockEnd:]
	fmt.Printf("Added import for %d/day%s\n", year, dayStr)

	// Write the updated content back to the file
	if err := os.WriteFile(mainFilePath, []byte(contentStr), 0644); err != nil {
		return fmt.Errorf("error writing main file: %w", err)
	}

	fmt.Printf("Updated main.go file\n")
	return nil
}


