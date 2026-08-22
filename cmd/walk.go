package cmd

import (
	"io/fs"
	"path/filepath"
)

type FileEntry struct {
	Path  string
	Size  int64
	Entry fs.DirEntry
}

func resolvePath(args []string) string {
	if len(args) == 1 {
		return args[0]
	}
	return "."
}

func walkFiles(root string) (files []FileEntry, dirs int, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// Path access failed
		if err != nil {
			return err
		}

		if d.Name() == ".git" {
			return fs.SkipDir
		}

		if d.IsDir() {
			dirs++
			return nil
		}

		fi, err := d.Info()
		if err != nil {
			return err
		}

		files = append(files, FileEntry{Path: path, Size: fi.Size(), Entry: d})
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return files, dirs, nil
}
