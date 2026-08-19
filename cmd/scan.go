package cmd

import (
	"fmt"
	"io/fs"
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

func scanfile_no_flag(args []string) error {
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
}

func scanfile_ext(args []string) error {
	var mpExt = make(map[string]int)

	for _, ext := range scanExtFilter {
		if ext[0] != '.' {
			ext = "." + ext
		}

		mpExt[ext] = 0
	}

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
		ext := filepath.Ext(name)

		if name == ".git" {
			return fs.SkipDir
		}

		if _, ok := mpExt[ext]; ok == false || d.IsDir() {
			return nil
		}

		mpExt[ext]++

		return nil
	})

	if err != nil {
		return err
	}

	for ext, num := range mpExt {
		ext = ext + ":"
		fmt.Printf("%-10s%d\n", ext, num)
	}

	return nil
}

func scanfile_top(args []string) error {
	nowPath := "."

	if len(args) == 1 {
		nowPath = args[0]
	}

	mpSize := make(map[string]int64)

	fileSlice := make([]string, 0)

	err := filepath.WalkDir(nowPath, func(path string, d fs.DirEntry, err error) error {
		// Path access failed
		if err != nil {
			return err
		}

		if d.Name() == ".git" {
			return fs.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		fileinfo, err := d.Info()

		if err != nil {
			return err
		}

		mpSize[path] = fileinfo.Size()
		fileSlice = append(fileSlice, path)
		return nil
	})

	scanTopSize = min(scanTopSize, int64(len(fileSlice)))

	if err != nil {
		return err
	}

	sort.Slice(fileSlice, func(i, j int) bool {
		if mpSize[fileSlice[i]] != mpSize[fileSlice[j]] {
			return mpSize[fileSlice[i]] > mpSize[fileSlice[j]]
		}
		return fileSlice[i] < fileSlice[j]
	})

	var maxLen int

	for i := 0; i < int(scanTopSize); i++ {
		maxLen = max(maxLen, runewidth.StringWidth(fileSlice[i]))
	}

	fmt.Printf("The top %d largest files:\n", scanTopSize)
	for i := 0; i < int(scanTopSize); i++ {
		fmt.Printf("%d. %s%s   %s\n",
			i+1,
			fileSlice[i],
			strings.Repeat(" ", maxLen-runewidth.StringWidth(fileSlice[i])),
			fileSizeJudge(mpSize[fileSlice[i]]))
	}

	return nil
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
		if len(scanExtFilter) != 0 {
			return scanfile_ext(args)
		} else if scanTopSize != 0 {
			return scanfile_top(args)
		} else {
			return scanfile_no_flag(args)
		}
	},

	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringSliceVarP(&scanExtFilter, "ext", "e", nil, "Filter by file extension")
	scanCmd.Flags().Int64VarP(&scanTopSize, "top", "t", 0, "Search for the top N largest files")
}
