```markdown
# Go Static Code Analysis String Search (GoSCASS)

`GoSCASS` is a command-line tool written in Go for performing advanced searches within source code files. It allows you to search for specific words or patterns across multiple files and directories, making it useful for static code analysis, code reviews, or finding instances of specific code patterns.

## Features

- **Recursive Directory Search**: Traverse directories recursively to search through all files.
- **File Type Filtering**: Limit the search to specific file types using the `-t` flag with support for multiple comma-separated extensions.
- **Custom Search Terms**: Provide custom search terms via the `-w` flag or use the embedded default list.
- **Regular Expression Support**: Use regular expressions for advanced search patterns with the `-r` flag.
- **Case Sensitivity**: Enable case-sensitive searches with the `-c` flag.
- **Context Lines**: Include a specified number of lines before and after each match using the `-n` flag for better context.
- **Output Formatting**: Results are written to a Markdown file with clickable links to the source files and clear indication of matched lines.
- **Default Search Terms**: If no search terms are provided, `GoSCASS` uses an embedded list of common indicators and potential issues for static code analysis.

## Installation

To build and install `GoSCASS`, you need to have Go installed on your system.

1. **Clone the Repository** (if applicable):

   ```bash
   git clone https://github.com/Sparagmos/GoSCASS.git
   cd GoSCASS
   ```

2. **Build the Executable**:

   ```bash
   go build -o goscass
   ```

3. **Install** (Optional):

   To install `GoSCASS` to your `$GOPATH/bin` directory:

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
        Comma-separated list of file extensions/types to limit the search to (e.g., .py,.ts,.json).
  -c    Enable case-sensitive search.
  -r    Interpret search words/strings as regular expressions.
  -n int
        Number of context lines to include before and after each match. (default 0)
  -o string
        Output file to write the results. (default "output.md")
```

### Notes

- **Default Search Terms**: If the `-w` flag is not provided, `GoSCASS` uses an embedded list of common search terms suitable for static code analysis, including indicators like `TODO`, `FIXME`, `BUG`, and potential security issues like `PASSWORD`, `SECRET`, `eval`, etc.
- **Multiple File Extensions**: The `-t` flag can accept multiple file extensions separated by commas (e.g., `.py,.ts,.json`) to search across different file types simultaneously.
- **Context Lines**: Use the `-n` flag to specify the number of lines before and after each match to include for better context in the output.

### Examples

#### 1. Search for Custom Terms in Python Files

```bash
goscass -w "password,secret" -t .py -o results.md
```

This command searches for the words `password` and `secret` in all `.py` files starting from the current directory and writes the results to `results.md`.

#### 2. Use Regular Expressions to Search in JavaScript and TypeScript Files

```bash
goscass -w "console\\.log|alert" -t .js,.ts -r -o js_ts_debug.md
```

This command searches for `console.log` or `alert` statements in all `.js` and `.ts` files using regular expressions and writes the results to `js_ts_debug.md`.

#### 3. Perform a Case-Sensitive Search for "FIXME" and "TODO" in All Files with Context Lines

```bash
goscass -w "FIXME,TODO" -c -n 1 -o fixme_todo.md
```

This command searches for `FIXME` and `TODO` in a case-sensitive manner in all files starting from the current directory, including one line before and after each match, and writes the results to `fixme_todo.md`.

#### 4. Use the Default Search Terms Across Multiple File Types with Context Lines

```bash
goscass -t .py,.ts,.json -n 2 -o code_analysis.md
```

This command searches using the embedded default search terms in `.py`, `.ts`, and `.json` files, including two lines before and after each match, and writes the results to `code_analysis.md`.

## Output Format

The output is written in Markdown format to the specified output file (default is `output.md`). Each match includes:

- A header with a clickable link to the file where the match was found, showing the relative path and line number.
- The matched line of code, along with specified context lines, enclosed in triple backticks for code formatting. The matched line is clearly indicated with a `>>` prefix.

**Example Output:**

```markdown
### [src/utils/helpers.py (Line 42)](src/utils/helpers.py)
```python
    40:     def some_function():
    41:         # Previous context line
>> 42:         # TODO: Refactor this function
    43:         def complex_function():
    44:             pass
```

### [flow/controllers/crm.controller.ts (Line 7)](flow/controllers/crm.controller.ts)
```typescript
    5:           await conn.login(
    6:             "salesforceadmin@sonosim.com",
>> 7:             "n3br4sk43030!" + "hngCeNEri8pHXl5Ie6LPSonN"
    8:           );
    9:         // Next context line
```
```

**Explanation:**

- **Header:** Shows the relative file path and the line number where the match was found, with a clickable link.
- **Code Block:**
  - Lines before and after the matched line provide context.
  - The matched line is prefixed with `>>` to clearly indicate the exact location of the match.
  - Language identifier (e.g., `python`, `typescript`) is used for syntax highlighting.

## Embedded Default Search Terms

The embedded default search terms include, but are not limited to:

- **General Indicators:** `TODO`, `FIXME`, `BUG`, `HACK`, `NOTE`, `OPTIMIZE`
- **Security-Related Terms:** `PASSWORD`, `SECRET`, `API_KEY`, `KEY`, `TOKEN`, `PRIVATE_KEY`, `PUBLIC_KEY`, `CREDENTIALS`
- **Suspicious Functions and Methods:**
  - **General:** `eval`, `exec`, `system`
  - **Python:** `pickle.loads`, `os.system`, `subprocess.Popen`, `input`, `paramiko`
  - **JavaScript:** `eval`, `document.write`, `innerHTML`, `console.log`, `alert`
  - **PHP:** `eval`, `exec`, `shell_exec`, `passthru`, `system`
  - **C/C++:** `gets`, `strcpy`, `strcat`, `sprintf`
- **Logging and Debugging:** `console.log`, `print`, `alert`
- **Potential Vulnerabilities:** `SELECT *`, `DROP TABLE`, `race condition`, `concurrency`
- **Suppressed Warnings:** `#nosec`, `@SuppressWarnings`
- **Licensing and Legal:** `LICENSE`, `COPYRIGHT`
- **Other Indicators:** `Not Implemented`, `TBD`, `Temporary`

## Limitations and Considerations

- **File Paths:** The program uses relative paths based on the specified root directory for generating links in the output Markdown file.
- **Permissions:** Ensure that you have read permissions for all files and directories you intend to search.
- **Binary Files:** The program does not filter out binary files. Use the `-t` flag to limit the search to specific file types.
- **Performance:** For very large codebases, the program may take some time to complete. Progress indicators are not provided.
- **Overlapping Context Lines:** When multiple matches are close to each other, context lines may overlap. Each match is treated independently.
- **False Positives:** Automated searches can produce false positives. Manually review the findings to assess the actual risk.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the tool.

## License

This project is licensed under the MIT License.

## Acknowledgments

- Thanks to all contributors and users who have provided feedback and suggestions.
- Inspired by common static code analysis tools and best practices in code security.
```
