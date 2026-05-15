package cmd

import (
	"os"

	lib "github.com/nelsonr/dbcook/internal"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Generates a database migration for a new entity table",
	Run:     runGenerateCmd,
}

func runGenerateCmd(cmd *cobra.Command, args []string) {
	name, err := lib.NormalizeName(args[0])
	if err != nil {
		cmd.PrintErrf("%s\n", err)
		os.Exit(1)
	}

	sql, err := lib.GenerateTableSql(name, args[1:])
	if err != nil {
		cmd.PrintErrf("%s\n", err)
		os.Exit(1)
	}

	filename, err := lib.GenerateFileName(name)
	if err != nil {
		cmd.PrintErrf("%s\n", err)
		os.Exit(1)
	}

	outputPath := lib.ResolveOutputPath(cmd, ".")
	err = lib.CreateFile(outputPath, filename, sql)
	if err != nil {
		cmd.PrintErrf("Error creating migration file: %s\n", filename)
		cmd.PrintErrf("%s\n", err)
		os.Exit(1)
	}

	cmd.Printf("Database migration file created successfully: %s\n", filename)
}

func init() {
	// Register the "generate" command
	//
	// Usage example:
	// dbcook generate posts title url
	//
	// This creates a migration file in the current working directory.
	rootCmd.AddCommand(generateCmd)

	// Register the "--output" flag for the "generate" command.
	//
	// Usage example:
	// dbcook generate posts title url --output db/migrations
	//
	// This creates a "db/migrations" output directory relative to the
	// current working directory for the generated migration file.
	generateCmd.Flags().String("output", ".", "output directory for the migration file")
}
