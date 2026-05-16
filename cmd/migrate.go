package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	lib "github.com/nelsonr/dbcook/internal"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Args:  cobra.MinimumNArgs(1),
	Short: "Runs existing database migrations",
	Run:   runMigrateCmd,
}

func runMigrateCmd(cmd *cobra.Command, args []string) {
	// 1. Check if database file exists
	dbPath := args[0]
	if !lib.FileExists(dbPath) {
		cmd.PrintErrf("Database file not found: %s\n", dbPath)
		os.Exit(1)
	}

	// 2. Open connection to database
	// The sqlite driver never errors on Open.
	db, _ := sql.Open("sqlite", dbPath)
	defer db.Close()
	fmt.Printf("Connected to database: %s\n\n", dbPath)

	// TODO: Add flag to override the migrations directory path
	//
	// 3. Retrieve the path to the migrations directory
	// If the config file is not present, it falls back to the
	// current working directory
	dirPath := "."
	conf, err := lib.ReadConfigFile(dirPath)
	if err == nil {
		dirPath = conf.OutputPath
	}

	// 4. Obtain list of migration filenames
	files, err := lib.ListMigrationFiles(dirPath)
	if err != nil {
		cmd.PrintErrf("Error reading files in %s", dirPath)
		os.Exit(1)
	}

	// Exit if no migration files were found
	if len(files) == 0 {
		if dirPath == "." {
			cmd.PrintErrln("No migration files found in the current directory")
		} else {
			cmd.PrintErrf("No migration files found in: %s\n", dirPath)
		}
		os.Exit(1)
	}

	// 5. Run migration files
	fmt.Printf("Found %d migration file(s)\n\n", len(files))

	for _, filename := range files {
		file, err := os.ReadFile(filepath.Join(dirPath, filename))
		if err != nil {
			cmd.PrintErrf("Error reading migration file: %s\n", filename)
			os.Exit(1)
		}

		fmt.Printf("Running migration: %s\n", filename)
		_, err = db.Exec(string(file))
		if err != nil {
			cmd.PrintErrf("Error while executing migration file, stopping...")
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}

	fmt.Printf("\nCompleted %d migrations successfully!\n", len(files))
}

func init() {
	// Register the "migrate" command
	//
	// Usage example:
	// dbcook migrate blog.db
	//
	// Runs the database migrations on the sqlite "blog.db" database
	rootCmd.AddCommand(migrateCmd)
}
