# Rovu

**Repository Overview Utility**

Rovu is a fast, lightweight command-line tool written in Go for analyzing local code repositories. It gives you a quick snapshot of a project's file count, directory structure, total size, and language distribution — right from your terminal.

## Features

- Recursively scan a directory, skipping `.git`
- Count files, directories, and total size
- Filter by file extension with `--ext`
- Show the top N largest files with `--top`
- Report file counts grouped by language/extension
- Simple, readable output with percentage bar charts

## Installation

Requires [Go 1.21+](https://go.dev/dl/).

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

Show the language distribution of a repository:

```bash
rovu stats .
rovu stats D:\project\medbot
```

If no path is given, the current directory (`.`) is scanned. The `.git` directory is always skipped.

### Example output

```
$ rovu scan .

Files:       11
Directories: 2
Size:        19.80 KB

$ rovu scan . --top 5

The top 5 largest files:
1. cmd\scan.go   3.90 KB
2. cmd\public.go 2.07 KB
3. README.md     2.05 KB
4. cmd\stats.go  2.01 KB

$ rovu stats .

File Extension Counts and Percentages (Ext | Num | Pcnt) :
Go                   │ 5          │  62.500% ████████████
Markdown             │ 2          │  25.000% █████
Git                  │ 1          │  12.500% ██
```

## Commands

| Command                    | Description                                           |
|----------------------------|-------------------------------------------------------|
| `rovu scan [path]`           | Count files, directories, and total size recursively |
| `rovu scan [path] --ext go,py` | Filter by file extension                            |
| `rovu scan [path] --top 10`   | Show the top N largest files                         |
| `rovu stats [path]`          | Count files by extension, with percentage and chart  |

## Roadmap

- `rovu todo [path]` — find TODO/FIXME/HACK comments with file and line numbers
- `rovu duplicate [path]` — find files with duplicate content
- `rovu scan . --json` — machine-readable output
- `rovu version` — print the version

## License

[MIT](LICENSE)
