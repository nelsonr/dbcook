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
	Run:     generateFn,
}

func generateFn(cmd *cobra.Command, args []string) {
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

	dirPath := "db/migrate"
	err = lib.CreateFile(dirPath, filename, sql)
	if err != nil {
		cmd.PrintErrf("Error creating sql file: %s\n", filename)
		cmd.PrintErrf("%s\n", err)
		os.Exit(1)
	}

	cmd.Printf("Database migration file created successfully: %s\n", filename)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
