package main

import (
	"fmt"
	"os"

	"github.com/AnasImloul/aoc-go/cmd/aoc-go/commands"
	"github.com/AnasImloul/aoc-go/pkg/utils"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:     "aoc-go",
	Short:   "Advent of Code CLI tool for Go",
	Long:    `A CLI tool for managing and running Advent of Code solutions in Go.`,
	Version: version,
}

func main() {
	// Add all commands
	rootCmd.AddCommand(
		commands.InitCmd,
		commands.GenerateCmd,
		commands.RunCmd,
		commands.TestCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		utils.Must(fmt.Fprintln(os.Stderr, err))
		os.Exit(1)
	}
}


