package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AnasImloul/aoc-go/pkg/generator"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate <year> <day>",
	Short: "Generate template files for a new day",
	Long:  `Generate template files for a new Advent of Code day including base, first, and second files.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
		day, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Invalid day: %v", err)
		}

		pattern, _ := cmd.Flags().GetString("pattern")
		withExample, _ := cmd.Flags().GetBool("example")

		if err := generator.GenerateFiles(year, day, pattern, withExample); err != nil {
			log.Fatalf("Error generating files: %v", err)
		}

		fmt.Printf("Generated files for year %d day %d\n", year, day)
		if pattern != "" {
			fmt.Printf("   Pattern: %s\n", pattern)
		}
		if withExample {
			fmt.Printf("   Generated example input file\n")
		}
	},
}

func init() {
	GenerateCmd.Flags().StringP("pattern", "p", "", "Template pattern (parser, grid, graph, simulation)")
	GenerateCmd.Flags().BoolP("example", "e", true, "Generate example input file")
}


