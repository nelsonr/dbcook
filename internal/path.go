package lib

import "github.com/spf13/cobra"

func ResolveOutputPath(cmd *cobra.Command, path string) string {
	conf, err := ReadConfigFile(path)
	if err == nil {
		path = conf.OutputPath
	}

	if cmd.Flags().Changed("output") {
		path, _ = cmd.Flags().GetString("output")
	}

	return path
}
