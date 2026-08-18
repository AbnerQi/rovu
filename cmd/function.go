package cmd

import (
	"fmt"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

var extDisplay = map[string]string{
	".go":     "Go",
	".py":     "Python",
	".java":   "Java",
	".rs":     "Rust",
	".c":      "C",
	".cpp":    "C++",
	".h":      "C Header",
	".hpp":    "C++ Header",
	".rb":     "Ruby",
	".php":    "PHP",
	".swift":  "Swift",
	".kt":     "Kotlin",
	".dart":   "Dart",
	".lua":    "Lua",
	".r":      "R",
	".jl":     "Julia",
	".zig":    "Zig",
	".nim":    "Nim",
	".cr":     "Crystal",
	".ex":     "Elixir",
	".exs":    "Elixir Script",
	".erl":    "Erlang",
	".hs":     "Haskell",
	".clj":    "Clojure",
	".scala":  "Scala",
	".groovy": "Groovy",

	".js":     "JavaScript",
	".ts":     "TypeScript",
	".jsx":    "React JSX",
	".tsx":    "React TSX",
	".html":   "HTML",
	".htm":    "HTML",
	".css":    "CSS",
	".scss":   "SCSS",
	".sass":   "SASS",
	".less":   "Less",
	".vue":    "Vue",
	".svelte": "Svelte",

	".json":     "JSON",
	".yaml":     "YAML",
	".yml":      "YAML",
	".toml":     "TOML",
	".xml":      "XML",
	".md":       "Markdown",
	".markdown": "Markdown",
	".sql":      "SQL",
	".graphql":  "GraphQL",
	".proto":    "Protobuf",

	".sh":   "Shell",
	".bash": "Bash",
	".zsh":  "Zsh",
	".fish": "Fish",
	".ps1":  "PowerShell",
	".bat":  "Batch",
	".cmd":  "Batch",

	".env":          "Env",
	".dockerignore": "Docker",
	".gitignore":    "Git",
	".npmrc":        "NPM",
	".yarnrc":       "Yarn",

	".txt":  "Text",
	".log":  "Log",
	".csv":  "CSV",
	".tsv":  "TSV",
	".ini":  "INI",
	".cfg":  "Config",
	".conf": "Config",

	"(none)": "No Extension",
}

func fileSizeJudge(size int64) string {
	switch {
	case size < KB:
		return fmt.Sprintf("%d bytes", size)
	case size < MB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	}
}

func barChart(count, total, bMax int) string {
	length_f := float64(count) / float64(total) * float64(bMax)
	length := int(length_f)

	if length == 0 {
		length = 1
	}

	return strings.Repeat("█", length)
}
