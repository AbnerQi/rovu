/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use: "scan",

	Short: "Scan a directory and report file, directory, and size statistics",

	Long: `Scan recursively walks a directory tree and reports the number
of files, directories, and total size, skipping version control
directories like .git.

If no path is given, the current directory is scanned.

Examples:

  rovu scan .
  rovu scan D:\project\rovu`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			fileNum int   // Total number of files
			dirNum  int   // Total number of directories
			sizeAll int64 // Total memory
		)

		nowPath := "."

		if len(args) == 1 {
			nowPath = args[0]
		}

		err := filepath.WalkDir(nowPath, func(_ string, d fs.DirEntry, err error) error {
			// Path access failed
			if err != nil {
				return err
			}

			name := d.Name()
			if name == ".git" {
				return fs.SkipDir
			}

			if d.IsDir() {
				dirNum++
			} else {
				fileNum++
				fileInfo, err := d.Info()

				// Failed to access file info
				if err != nil {
					return err
				}

				sizeAll = sizeAll + fileInfo.Size()
			}

			return nil
		})

		if err != nil {
			return err
		}

		fmt.Printf("%-12s %d\n", "Files:", fileNum)
		fmt.Printf("%-12s %d\n", "Directories:", dirNum)
		fmt.Printf("%-12s %s\n", "Size:", fileSizeJudge(sizeAll))

		return nil
	},

	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
