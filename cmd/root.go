package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RovuCmd = &cobra.Command{
	Use: "rovu",

	Short: "A fast local repository analysis CLI",

	Long: `Rovu is a fast local repository analysis CLI.
It scans a code repository and reports useful statistics such as
file counts, directory counts, total size, language distribution,
and the largest files.

Examples:

  rovu scan .
  rovu stats .`,
}

func Execute() {
	err := RovuCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
