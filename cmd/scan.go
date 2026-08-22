package cmd

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
)

var (
	scanExtFilter []string // Filter by file extension
	scanTopSize   int64    // Search for the top N largest files
)

func filterByExt(files []FileEntry, exts []string) ([]FileEntry, error) {
	allowed := make(map[string]bool)
	for _, ext := range exts {
		if ext == "" {
			continue
		}

		if ext[0] != '.' {
			ext = "." + ext
		}

		allowed[ext] = true
	}

	if len(allowed) == 0 {
		return nil, fmt.Errorf("no valid extension provided to --ext")
	}

	filtered := make([]FileEntry, 0, len(files))
	for _, f := range files {
		if allowed[filepath.Ext(f.Path)] {
			filtered = append(filtered, f)
		}
	}

	return filtered, nil
}

func largestFiles(files []FileEntry, n int64) []FileEntry {
	if int64(len(files)) < n {
		n = int64(len(files))
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].Size != files[j].Size {
			return files[i].Size > files[j].Size
		}
		return files[i].Path < files[j].Path
	})

	return files[:n]
}

func printCounts(files []FileEntry, dirs int) {
	var fileNum int
	var sizeAll int64
	for _, f := range files {
		fileNum++
		sizeAll += f.Size
	}

	fmt.Printf("%-12s %d\n", "Files:", fileNum)
	fmt.Printf("%-12s %d\n", "Directories:", dirs)
	fmt.Printf("%-12s %s\n", "Size:", formatSize(sizeAll))
}

func printExtCounts(files []FileEntry) {
	counts := make(map[string]int)
	for _, f := range files {
		counts[filepath.Ext(f.Path)]++
	}

	exts := make([]string, 0, len(counts))
	for ext := range counts {
		exts = append(exts, ext)
	}

	sort.Slice(exts, func(i, j int) bool {
		if counts[exts[i]] == counts[exts[j]] {
			return exts[i] < exts[j]
		}
		return counts[exts[i]] > counts[exts[j]]
	})

	for _, ext := range exts {
		fmt.Printf("%-10s%d\n", ext+":", counts[ext])
	}
}

func printTopFiles(files []FileEntry) {
	var maxLen int
	for _, f := range files {
		maxLen = max(maxLen, runewidth.StringWidth(f.Path))
	}

	fmt.Printf("The top %d largest files:\n", len(files))
	for i, f := range files {
		fmt.Printf("%d. %s%s   %s\n",
			i+1,
			f.Path,
			strings.Repeat(" ", maxLen-runewidth.StringWidth(f.Path)),
			formatSize(f.Size))
	}
}

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
		if scanTopSize < 0 {
			return fmt.Errorf("invalid --top value %d: must be a non-negative integer", scanTopSize)
		}

		files, dirs, err := walkFiles(resolvePath(args))
		if err != nil {
			return err
		}

		useExt := len(scanExtFilter) != 0
		useTop := scanTopSize != 0

		if useExt {
			files, err = filterByExt(files, scanExtFilter)
			if err != nil {
				return err
			}
		}

		if useTop {
			files = largestFiles(files, scanTopSize)
		}

		switch {
		case useTop:
			printTopFiles(files)
		case useExt:
			printExtCounts(files)
		default:
			printCounts(files, dirs)
		}

		return nil
	},

	Args: cobra.MaximumNArgs(1),
}

func init() {
	RovuCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringSliceVarP(&scanExtFilter, "ext", "e", nil, "Filter by file extension")
	scanCmd.Flags().Int64VarP(&scanTopSize, "top", "t", 0, "Search for the top N largest files")
}
