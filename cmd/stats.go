/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show language distribution of mainstream programming languages",
	Long: `Stats walks a directory and tallies files by their extension,
reporting only mainstream programming languages and common
config/document formats. Unrecognized or niche extensions are
not listed.

Results are sorted by file count, from most to least common.

If no path is given, the current directory is scanned.

Examples:

  rovu stats .
  rovu stats D:\project\bot
  rovu stats . --top 10`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nowPath := "."

		if len(args) == 1 {
			nowPath = args[0]
		}

		extMp := make(map[string]int)

		err := filepath.WalkDir(nowPath, func(path string, d fs.DirEntry, err error) error {
			// Path access failed
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			name := d.Name()
			ext := filepath.Ext(name)

			if ext == "" || extDisplay[ext] == "" {
				return nil
			}

			extMp[ext]++

			return nil
		})

		// extMp save the number of each "ext"
		if err != nil {
			return err
		}

		if len(extMp) == 0 {
			println("No file with extension!")
			return nil
		}

		exts := make([]string, 0, len(extMp))
		for ext := range extMp {
			exts = append(exts, ext)
		}

		sort.Slice(exts, func(i, j int) bool {
			if extMp[exts[i]] == extMp[exts[j]] {
				return exts[i] < exts[j]
			}
			return extMp[exts[i]] > extMp[exts[j]]
		})

		var total int
		for _, ext := range exts {
			total += extMp[ext]
		}

		fmt.Println("File Extension Counts and Percentages (Ext | Num | Pcnt) :")
		for _, ext := range exts {
			pcnt := float64(extMp[ext]) / float64(total) * 100
			bar := barChart(extMp[ext], total, 20)
			fmt.Printf("%-20s │ %-10d │ %7.3f%% %s\n", extDisplay[ext], extMp[ext], pcnt, bar)
		}

		return nil
	},
}

func init() {
	RovuCmd.AddCommand(statsCmd)
}
