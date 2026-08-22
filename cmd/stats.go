package cmd

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

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
  rovu stats D:\project\bot`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		files, _, err := walkFiles(resolvePath(args))
		if err != nil {
			return err
		}

		extCounts := make(map[string]int)
		for _, f := range files {
			ext := filepath.Ext(f.Path)
			if ext == "" || extDisplay[ext] == "" {
				continue
			}

			extCounts[ext]++
		}

		if len(extCounts) == 0 {
			fmt.Println("No file with extension!")
			return nil
		}

		exts := make([]string, 0, len(extCounts))
		for ext := range extCounts {
			exts = append(exts, ext)
		}

		sort.Slice(exts, func(i, j int) bool {
			if extCounts[exts[i]] == extCounts[exts[j]] {
				return exts[i] < exts[j]
			}
			return extCounts[exts[i]] > extCounts[exts[j]]
		})

		var total int
		for _, ext := range exts {
			total += extCounts[ext]
		}

		fmt.Println("File Extension Counts and Percentages (Ext | Num | Pcnt) :")
		for _, ext := range exts {
			pcnt := float64(extCounts[ext]) / float64(total) * 100
			bar := barChart(extCounts[ext], total, 20)
			fmt.Printf("%-20s │ %-10d │ %7.3f%% %s\n", extDisplay[ext], extCounts[ext], pcnt, bar)
		}

		return nil
	},
}

func init() {
	RovuCmd.AddCommand(statsCmd)
}
