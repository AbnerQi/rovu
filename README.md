# Rovu

**Repository Overview Utility**

Rovu is a fast, lightweight command-line tool written in Go for analyzing local code repositories. It gives you a quick snapshot of a project's file count, directory structure, total size, and language distribution — right from your terminal.

## Features

- Recursively scan a directory, skipping `.git`
- Count files, directories, and total size
- Filter by file extension with `--ext`
- Show the top N largest files with `--top`
- Combine filters: `scan . --ext go --top 10` lists the 10 largest Go files
- Report file counts grouped by language/extension
- Simple, readable output with percentage bar charts

## Installation

Requires [Go 1.26+](https://go.dev/dl/).

```bash
go install github.com/AbnerQi/rovu@latest
```

Or build from source:

```bash
git clone https://github.com/AbnerQi/rovu.git
cd rovu
go build -o rovu .
```

## Usage

Scan a repository and report file, directory, and size statistics:

```bash
rovu scan .
rovu scan D:\project\medbot
```

Filter by file extension:

```bash
rovu scan . --ext go
rovu scan . --ext go,py
```

Show the top N largest files:

```bash
rovu scan . --top 10
```

Combine both filters — the 10 largest Go files:

```bash
rovu scan . --ext go --top 10
```

Show the language distribution of a repository:

```bash
rovu stats .
rovu stats D:\project\medbot
```

If no path is given, the current directory (`.`) is scanned. The `.git` directory is always skipped.

### Example output

```
$ rovu scan .

Files:       13
Directories: 2
Size:        24.44 KB

$ rovu scan . --ext go --top 3

The top 3 largest files:
1. cmd\scan.go         3.44 KB
2. cmd\stats.go        1.72 KB
3. cmd\extensions.go   1.45 KB

$ rovu stats .

File Extension Counts and Percentages (Ext | Num | Pcnt) :
Go                   │ 7          │  70.000% ██████████████
Markdown             │ 2          │  20.000% ████
Git                  │ 1          │  10.000% ██
```

## Commands

| Command                          | Description                                          |
|----------------------------------|------------------------------------------------------|
| `rovu scan [path]`               | Count files, directories, and total size recursively |
| `rovu scan [path] --ext go,py`   | Filter by file extension                             |
| `rovu scan [path] --top 10`      | Show the top N largest files                         |
| `rovu scan [path] --ext go --top 10` | Top N largest files of a given extension         |
| `rovu stats [path]`              | Count files by extension, with percentage and chart  |

## Embed as a library

`RovuCmd` is exported so other Cobra-based tools can embed rovu as a subcommand. For example, the `z` personal CLI toolkit wires it in:

```go
import rovucmd "github.com/AbnerQi/rovu/cmd"

func init() {
	rootCmd.AddCommand(rovucmd.RovuCmd)
}
```

This gives `z rovu scan .` — the full rovu command tree under `z`, while `rovu` keeps working as a standalone binary.

## Roadmap

- `rovu loc [path]` — count code/blank/comment lines by language
- `rovu secret [path]` — scan for leaked API keys, passwords, and tokens
- `rovu activity [path]` — show git commit and contributor activity
- `rovu health [path]` — score a repository's health (README, tests, CI, large files)
- `rovu duplicate [path]` — find files with duplicate content
- `rovu version` — print the version

## License

[MIT](LICENSE)
