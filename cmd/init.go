package cmd

import (
	"strings"

	lib "github.com/nelsonr/dbcook/internal"
	"github.com/spf13/cobra"
)

const configFilename string = "dbcook.toml"

const configContent string = `
# Sets the output path for migration files
output_path = "db/migrations"
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes dbcook.toml config file",
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	lib.CreateFile(".", configFilename, strings.TrimSpace(configContent)+"\n")
}

func init() {
	// Register the "init" command
	//
	// Usage example:
	// dbcook init
	//
	// This creates a config file named "dbcook.toml"
	// in the current working directory.
	rootCmd.AddCommand(initCmd)
}
