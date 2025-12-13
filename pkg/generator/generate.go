package generator

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/AnasImloul/aoc-go/internal/constants"
	"github.com/AnasImloul/aoc-go/pkg/input"
)

// Templates holds embedded template files
//
//go:embed templates/*.tmpl
var Templates embed.FS

// transaction tracks all changes made during generation for rollback
type transaction struct {
	createdFiles  []string
	createdDirs   []string
	modifiedFiles map[string][]byte // original content for rollback
}

func newTransaction() *transaction {
	return &transaction{
		createdFiles:  make([]string, 0),
		createdDirs:   make([]string, 0),
		modifiedFiles: make(map[string][]byte),
	}
}

// rollback undoes all changes made during the transaction.
// Errors during rollback are logged but do not prevent cleanup from continuing.
func (t *transaction) rollback() {
	if len(t.createdFiles) == 0 && len(t.modifiedFiles) == 0 && len(t.createdDirs) == 0 {
		return
	}

	fmt.Println("\nRolling back changes...")

	// Remove created files
	for _, f := range t.createdFiles {
		if err := os.Remove(f); err != nil {
			fmt.Fprintf(os.Stderr, "  Warning: failed to remove %s: %v\n", f, err)
		} else {
			fmt.Printf("  Removed: %s\n", f)
		}
	}

	// Restore modified files
	for path, content := range t.modifiedFiles {
		if err := os.WriteFile(path, content, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "  Warning: failed to restore %s: %v\n", path, err)
		} else {
			fmt.Printf("  Restored: %s\n", path)
		}
	}

	// Remove created directories (in reverse order to handle nested dirs)
	for i := len(t.createdDirs) - 1; i >= 0; i-- {
		if err := os.Remove(t.createdDirs[i]); err != nil {
			fmt.Fprintf(os.Stderr, "  Warning: failed to remove directory %s: %v\n", t.createdDirs[i], err)
		} else {
			fmt.Printf("  Removed directory: %s\n", t.createdDirs[i])
		}
	}
}

// createDir creates a directory and tracks it for potential rollback
func (t *transaction) createDir(path string) error {
	// Check if directory already exists
	if _, err := os.Stat(path); err == nil {
		return nil // Already exists, nothing to track
	}

	// Find the first non-existent parent to track
	toCreate := []string{}
	current := path
	for {
		if _, err := os.Stat(current); err == nil {
			break
		}
		toCreate = append([]string{current}, toCreate...)
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	t.createdDirs = append(t.createdDirs, toCreate...)
	return nil
}

// createFile creates a file and tracks it for potential rollback
func (t *transaction) createFile(path string, content []byte) error {
	// Check if file already exists
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file already exists: %s", path)
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	t.createdFiles = append(t.createdFiles, path)
	return nil
}

// modifyFile modifies a file and tracks original content for rollback
func (t *transaction) modifyFile(path string, newContent []byte) error {
	// Read original content for rollback
	original, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Only track if we haven't already
	if _, exists := t.modifiedFiles[path]; !exists {
		t.modifiedFiles[path] = original
	}

	return os.WriteFile(path, newContent, 0644)
}

// GenerateFiles generates the solution files for a given year and day.
func GenerateFiles(year int, day int, pattern string, withExample bool) error {
	tx := newTransaction()
	defer func() {
		// Only rollback if there was an error (tx will track if commit needed)
		if r := recover(); r != nil {
			tx.rollback()
			panic(r)
		}
	}()

	yearStr := strconv.Itoa(year)
	dayStr := fmt.Sprintf("%02d", day)
	basePath := filepath.Join(".", constants.SolutionsDir, yearStr, fmt.Sprintf(constants.DayFolderName, dayStr))

	moduleName, err := getModuleName()
	if err != nil {
		tx.rollback()
		return fmt.Errorf("error reading module name: %w", err)
	}

	if err := tx.createDir(basePath); err != nil {
		tx.rollback()
		return fmt.Errorf("error creating directory: %w", err)
	}

	solutionTemplate := determineSolutionTemplate(pattern)
	if err := generateSolutionFiles(tx, basePath, solutionTemplate, moduleName, year, day, yearStr, dayStr); err != nil {
		tx.rollback()
		return err
	}

	if err := updateMainFileTransactional(tx, moduleName, year, day); err != nil {
		tx.rollback()
		return fmt.Errorf("error updating main file: %w", err)
	}

	if withExample {
		if err := createExampleFilesTransactional(tx, year, day); err != nil {
			tx.rollback()
			return fmt.Errorf("error creating example files: %w", err)
		}
	}

	if _, err := input.TryRead(year, day); err != nil {
		tx.rollback()
		return fmt.Errorf("error fetching input: %w", err)
	}

	fmt.Printf("Folder and files created successfully at %s\n", basePath)
	return nil
}

func determineSolutionTemplate(pattern string) string {
	if pattern == "" {
		return "solution.tmpl"
	}

	patternTemplate := fmt.Sprintf("base_%s.tmpl", pattern)
	if _, err := Templates.ReadFile("templates/" + patternTemplate); err == nil {
		return patternTemplate
	}

	fmt.Printf("Pattern template %s not found, using default solution.tmpl\n", patternTemplate)
	return "solution.tmpl"
}

func generateSolutionFiles(tx *transaction, basePath, solutionTemplate, moduleName string, year, day int, yearStr, dayStr string) error {
	files := []struct {
		TemplateFile string
		OutputFile   string
	}{
		{solutionTemplate, "solution.go"},
		{"part1.tmpl", "part1.go"},
		{"part2.tmpl", "part2.go"},
	}

	for _, file := range files {
		outputFilePath := filepath.Join(basePath, file.OutputFile)

		if _, err := os.Stat(outputFilePath); err == nil {
			fmt.Printf("File %s already exists, skipping...\n", outputFilePath)
			continue
		}

		if err := createFileFromTemplate(tx, file.TemplateFile, outputFilePath, moduleName, year, day, yearStr, dayStr); err != nil {
			return fmt.Errorf("error creating file %s: %w", outputFilePath, err)
		}
	}
	return nil
}

func createFileFromTemplate(tx *transaction, templateFile, outputPath, moduleName string, year, day int, yearStr, dayStr string) error {
	templateContent, err := Templates.ReadFile("templates/" + templateFile)
	if err != nil {
		return fmt.Errorf("error reading template file %s: %w", templateFile, err)
	}

	tmpl, err := template.New(templateFile).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template %s: %w", templateFile, err)
	}

	templateData := struct {
		ModuleName string
		Year       int
		Day        int
		YearStr    string
		DayStr     string
		DayPackage string
	}{
		ModuleName: moduleName,
		Year:       year,
		Day:        day,
		YearStr:    yearStr,
		DayStr:     dayStr,
		DayPackage: fmt.Sprintf("day%s", dayStr),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
		return fmt.Errorf("error executing template %s: %w", templateFile, err)
	}

	if err := tx.createFile(outputPath, buf.Bytes()); err != nil {
		return fmt.Errorf("error creating file %s: %w", outputPath, err)
	}

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

func createExampleFilesTransactional(tx *transaction, year, day int) error {
	dayStr := fmt.Sprintf("%02d", day)
	dayFolderName := fmt.Sprintf("day%s", dayStr)

	// Create the day-specific folder: data/examples/2025/day01/
	examplesPath := filepath.Join(".", constants.DataDir, constants.ExamplesDir, strconv.Itoa(year), dayFolderName)

	// Create the directory structure
	if err := tx.createDir(examplesPath); err != nil {
		return fmt.Errorf("error creating examples directory: %w", err)
	}

	// Create example input file: data/examples/2025/day01/input.txt
	inputFile := filepath.Join(examplesPath, constants.ExampleInputFileName)
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		content := "# Paste example input here\n"
		if err := tx.createFile(inputFile, []byte(content)); err != nil {
			return fmt.Errorf("error creating example input file: %w", err)
		}
		fmt.Printf("Created example input file: %s\n", inputFile)
	}

	// Create expected output files for part 1 and part 2: data/examples/2025/day01/part1.txt, part2.txt
	for _, partNum := range []string{"1", "2"} {
		outputFile := filepath.Join(examplesPath, fmt.Sprintf(constants.PartFileName, partNum))
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			content := "# Expected output for part " + partNum + "\n"
			if err := tx.createFile(outputFile, []byte(content)); err != nil {
				return fmt.Errorf("error creating expected output file: %w", err)
			}
			fmt.Printf("Created expected output file: %s\n", outputFile)
		}
	}

	return nil
}

func updateMainFileTransactional(tx *transaction, moduleName string, year, day int) error {
	dayStr := fmt.Sprintf("%02d", day)
	mainFilePath := constants.MainGoFile

	if _, err := os.Stat(mainFilePath); err != nil {
		return fmt.Errorf("%s file does not exist: %w", mainFilePath, err)
	}

	content, err := os.ReadFile(mainFilePath)
	if err != nil {
		return fmt.Errorf("error reading main file: %w", err)
	}

	importStatement := fmt.Sprintf("_ \"%s/solutions/%d/day%s\"", moduleName, year, dayStr)
	if importExists(string(content), importStatement) {
		fmt.Printf("Import for %d/day%s already exists, skipping...\n", year, dayStr)
		return nil
	}

	updatedContent, err := insertImport(string(content), importStatement)
	if err != nil {
		return fmt.Errorf("error inserting import: %w", err)
	}

	if err := tx.modifyFile(mainFilePath, []byte(updatedContent)); err != nil {
		return fmt.Errorf("error writing main file: %w", err)
	}

	fmt.Printf("Added import for %d/day%s\n", year, dayStr)
	fmt.Printf("Updated %s file\n", mainFilePath)
	return nil
}

func importExists(content, importStatement string) bool {
	importPattern := regexp.MustCompile(regexp.QuoteMeta(importStatement))
	return importPattern.MatchString(content)
}

func insertImport(content, importStatement string) (string, error) {
	importStart := strings.Index(content, "import (")
	if importStart == -1 {
		return "", fmt.Errorf("could not find import block in main.go")
	}

	importBlockEnd := findImportBlockEnd(content, importStart)
	if importBlockEnd == -1 {
		return "", fmt.Errorf("could not find end of import block in main.go")
	}

	// Insert the new import before the closing parenthesis
	updated := content[:importBlockEnd] + "\t" + importStatement + "\n" + content[importBlockEnd:]
	return updated, nil
}

func findImportBlockEnd(content string, importStart int) int {
	searchStart := importStart + len("import (")
	depth := 1
	for i := searchStart; i < len(content); i++ {
		if content[i] == '(' {
			depth++
		} else if content[i] == ')' {
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}
