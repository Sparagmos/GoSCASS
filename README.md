# Go Static Code Analysis String Search (goscass)

`goscass` is a command-line tool written in Go for performing advanced searches within source code files. It allows you to search for specific words or patterns across multiple files and directories, making it useful for static code analysis, code reviews, or finding instances of specific code patterns.

## Features

- **Recursive Directory Search**: Traverse directories recursively to search through all files.
- **File Type Filtering**: Limit the search to specific file types using the `-t` flag.
- **Custom Search Terms**: Provide custom search terms via the `-w` flag or use the embedded default list.
- **Regular Expression Support**: Use regular expressions for advanced search patterns with the `-r` flag.
- **Case Sensitivity**: Enable case-sensitive searches with the `-c` flag.
- **Output Formatting**: Results are written to a Markdown file with clickable links to the source files.
- **Default Search Terms**: If no search terms are provided, `goscass` uses an embedded list of common indicators and potential issues for static code analysis.

## Installation

To build and install `goscass`, you need to have Go installed on your system.

1. **Clone the Repository** (if applicable):

   ```bash
   git clone https://github.com/yourusername/goscass.git
   cd goscass
   ```

2. **Build the Executable**:

   ```bash
   go build -o goscass
   ```

3. **Install** (Optional):

   To install `goscass` to your `$GOPATH/bin` directory:

   ```bash
   go install
   ```

## Usage

```bash
goscass [options]
```

### Command-Line Options

```
Usage of goscass:
  -w string
        Comma-separated list of words/strings to search for, or a file containing words (one per line).
  -d string
        Directory to start the search in. (default ".")
  -t string
        File extension/type to limit the search to (e.g., .py).
  -c    Enable case-sensitive search.
  -r    Interpret search words/strings as regular expressions.
  -o string
        Output file to write the results. (default "output.md")
```

### Notes

- **Default Search Terms**: If the `-w` flag is not provided, `goscass` uses an embedded list of common search terms suitable for static code analysis, including indicators like `TODO`, `FIXME`, `BUG`, and potential security issues like `PASSWORD`, `SECRET`, `eval`, etc.

### Examples

#### 1. Search for Custom Terms in Python Files

```bash
goscass -w "password,secret" -t .py -o results.md
```

This command searches for the words `password` and `secret` in all `.py` files starting from the current directory and writes the results to `results.md`.

#### 2. Use Regular Expressions to Search in JavaScript Files

```bash
goscass -w "console\\.log|alert" -t .js -r -o js_debug.md
```

This command searches for `console.log` or `alert` statements in all `.js` files using regular expressions and writes the results to `js_debug.md`.

#### 3. Perform a Case-Sensitive Search for "FIXME" and "TODO" in All Files

```bash
goscass -w "FIXME,TODO" -c -o fixme_todo.md
```

This command searches for `FIXME` and `TODO` in a case-sensitive manner in all files starting from the current directory.

#### 4. Use the Default Search Terms

```bash
goscass -o code_analysis.md
```

This command searches using the embedded default search terms and writes the results to `code_analysis.md`.

## Output Format

The output is written in Markdown format to the specified output file (default is `output.md`). Each match includes:

- A header with a clickable link to the file where the match was found, showing the relative path and line number.
- The matched line of code enclosed in triple backticks for code formatting.

**Example Output:**

```markdown
### [src/utils/helpers.py (Line 42)](src/utils/helpers.py)
```

```python
# TODO: Refactor this function
def complex_function():
    pass
```

## Embedded Default Search Terms

The embedded default search terms include, but are not limited to:

- **General Indicators**: `TODO`, `FIXME`, `BUG`, `HACK`, `NOTE`, `OPTIMIZE`
- **Security-Related Terms**: `PASSWORD`, `SECRET`, `API_KEY`, `KEY`, `TOKEN`, `PRIVATE_KEY`, `PUBLIC_KEY`, `CREDENTIALS`
- **Suspicious Functions and Methods**:
  - **General**: `eval`, `exec`, `system`
  - **Python**: `pickle.loads`, `os.system`, `subprocess.Popen`, `input`, `paramiko`
  - **JavaScript**: `eval`, `document.write`, `innerHTML`, `console.log`, `alert`
  - **PHP**: `eval`, `exec`, `shell_exec`, `passthru`, `system`
  - **C/C++**: `gets`, `strcpy`, `strcat`, `sprintf`
- **Logging and Debugging**: `console.log`, `print`, `alert`
- **Potential Vulnerabilities**: `SELECT *`, `DROP TABLE`, `race condition`, `concurrency`
- **Suppressed Warnings**: `#nosec`, `@SuppressWarnings`
- **Licensing and Legal**: `LICENSE`, `COPYRIGHT`
- **Other Indicators**: `Not Implemented`, `TBD`, `Temporary`

## Limitations and Considerations

- **File Paths**: The program uses relative paths based on the specified root directory for generating links in the output Markdown file.
- **Permissions**: Ensure that you have read permissions for all files and directories you intend to search.
- **Binary Files**: The program does not filter out binary files. Use the `-t` flag to limit the search to specific file types.
- **Performance**: For very large codebases, the program may take some time to complete. Progress indicators are not provided.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the tool.

## License

This project is licensed under the MIT License.

## Acknowledgments

- Thanks to all contributors and users who have provided feedback and suggestions.
