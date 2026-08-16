package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "rovu",

	Short: "A fast local repository analysis CLI",

	Long: `Rovu is a fast local repoitory analysis CLI.
	It scans a code repository and reports useful statistics such as
	file counts, directory counts, total size, language distribution,
	large files, TODO/FIXME comments, and duplicate files.

	Examples:

	rovu scan .`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
