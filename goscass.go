package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type SearchOptions struct {
	Words         []string
	Directory     string
	FileTypes     []string // Changed to a slice to support multiple file types
	CaseSensitive bool
	UseRegex      bool
	OutputFile    string
	ContextLines  int // New field for context lines
}

type Match struct {
	FilePath string
	LineNum  int
	Line     string
	Context  []string // New field for context lines
}

// **Embedded Default Search Terms**
var defaultSearchTerms = []string{
	"TODO",
	"FIXME",
	"BUG",
	"HACK",
	"PASSWORD",
	"SECRET",
	"API_KEY",
	"KEY",
	"TOKEN",
	"PRIVATE_KEY",
	"PUBLIC_KEY",
	"CREDENTIALS",
	"eval",
	"exec",
	"system",
	"pickle.loads",
	"os.system",
	"subprocess.Popen",
	"input",
	"paramiko",
	"document.write",
	"innerHTML",
	"console.log",
	"print",
	"alert",
	"SELECT *",
	"DROP TABLE",
	"race condition",
	"concurrency",
	"#nosec",
	"@SuppressWarnings",
	"Not Implemented",
	"TBD",
	"Temporary",
}

func main() {
	// Parse command-line arguments
	options := parseFlags()

	fmt.Printf("Search terms: %v\n", options.Words)
	fmt.Printf("File types: %v\n", options.FileTypes)
	fmt.Printf("Context lines: %d\n", options.ContextLines)

	// Prepare words/regex patterns
	patterns := preparePatterns(options.Words, options.CaseSensitive, options.UseRegex)

	// Channel to collect matches
	matchesChan := make(chan Match, 100)

	// WaitGroup to wait for all searchFile goroutines to finish
	var wg sync.WaitGroup

	// Start walking the directory tree
	err := filepath.Walk(options.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accessing path %q: %v\n", path, err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check file extensions
		if len(options.FileTypes) > 0 {
			if !isFileType(path, options.FileTypes) {
				return nil
			}
		}

		// Proceed to search the file
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			searchFile(p, patterns, options.CaseSensitive, matchesChan, options.ContextLines)
		}(path)

		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking the path %q: %v\n", options.Directory, err)
		os.Exit(1)
	}

	// Wait for all searchFile goroutines to finish and then close matchesChan
	go func() {
		wg.Wait()
		close(matchesChan)
	}()

	// Collect matches and write to output file
	writeMatches(matchesChan, options.OutputFile, options.Directory, options.FileTypes)
}

func parseFlags() SearchOptions {
	wordsPtr := flag.String("w", "", "Comma-separated list of words/strings to search for, or a file containing words (one per line).")
	dirPtr := flag.String("d", ".", "Directory to start the search in.")
	typePtr := flag.String("t", "", "Comma-separated list of file extensions/types to limit the search to (e.g., .py,.ts,.json).")
	casePtr := flag.Bool("c", false, "Enable case-sensitive search.")
	regexPtr := flag.Bool("r", false, "Interpret search words/strings as regular expressions.")
	outputPtr := flag.String("o", "output.md", "Output file to write the results.")
	contextPtr := flag.Int("n", 0, "Number of context lines to include in the output.")

	flag.Parse()

	var words []string

	if *wordsPtr != "" {
		if strings.Contains(*wordsPtr, ",") {
			// Comma-separated list
			words = strings.Split(*wordsPtr, ",")
		} else if fileExists(*wordsPtr) {
			// Read words from file
			fileWords, err := readWordsFromFile(*wordsPtr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading words from file: %v\n", err)
				os.Exit(1)
			}
			words = fileWords
		} else {
			// Single word
			words = []string{*wordsPtr}
		}
	} else {
		// Use embedded default search terms
		words = defaultSearchTerms
	}

	var fileTypes []string
	if *typePtr != "" {
		fileTypes = strings.Split(*typePtr, ",")
	}

	return SearchOptions{
		Words:         words,
		Directory:     *dirPtr,
		FileTypes:     fileTypes, // Now a slice to support multiple types
		CaseSensitive: *casePtr,
		UseRegex:      *regexPtr,
		OutputFile:    *outputPtr,
		ContextLines:  *contextPtr,
	}
}

func preparePatterns(words []string, caseSensitive bool, useRegex bool) []*regexp.Regexp {
	patterns := make([]*regexp.Regexp, 0, len(words))
	for _, word := range words {
		var pattern string
		if useRegex {
			pattern = word
		} else {
			pattern = regexp.QuoteMeta(word)
		}

		var re *regexp.Regexp
		var err error
		if caseSensitive {
			re, err = regexp.Compile(pattern)
		} else {
			re, err = regexp.Compile("(?i)" + pattern)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid pattern %q: %v\n", pattern, err)
		} else {
			fmt.Printf("Compiled pattern: %s\n", re.String())
			patterns = append(patterns, re)
		}
	}
	return patterns
}

func searchFile(path string, patterns []*regexp.Regexp, caseSensitive bool, matchesChan chan<- Match, contextLines int) {
	fmt.Printf("Searching file: %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %q: %v\n", path, err)
		return
	}
	defer file.Close()

	// Read all lines into a slice
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", path, err)
		return
	}

	// Iterate over lines and search for matches
	for i, line := range lines {
		for _, re := range patterns {
			if re.MatchString(line) {
				fmt.Printf("Match found in %s at line %d: %s\n", path, i+1, line)

				// Collect context lines
				start := i - contextLines
				if start < 0 {
					start = 0
				}
				end := i + contextLines + 1
				if end > len(lines) {
					end = len(lines)
				}
				context := lines[start:end]

				matchesChan <- Match{
					FilePath: path,
					LineNum:  i + 1,
					Line:     line,
					Context:  context,
				}
				break
			}
		}
	}
}

func writeMatches(matchesChan <-chan Match, outputFile string, rootDir string, fileTypes []string) {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file %q: %v\n", outputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for match := range matchesChan {
		// Compute the relative file path
		relPath, err := filepath.Rel(rootDir, match.FilePath)
		if err != nil {
			relPath = match.FilePath // Fallback to absolute path if error occurs
		}

		// Format the header with a Markdown link
		displayText := fmt.Sprintf("%s (Line %d)", relPath, match.LineNum)
		link := relPath
		header := fmt.Sprintf("### [%s](%s)\n", displayText, link)

		// Write to the output file
		writer.WriteString(header)

		// Determine language identifier based on file extension
		languageID := strings.TrimPrefix(filepath.Ext(match.FilePath), ".")

		if languageID != "" {
			writer.WriteString(fmt.Sprintf("```%s\n", languageID))
		} else {
			writer.WriteString("```\n")
		}

		// Write context lines
		for i, ctxLine := range match.Context {
			ctxLineNum := match.LineNum - len(match.Context)/2 + i
			if ctxLineNum < 1 {
				ctxLineNum = 1
			}
			prefix := "    "
			if ctxLineNum == match.LineNum {
				prefix = ">>  " // Indicate the matched line
			}
			writer.WriteString(fmt.Sprintf("%s%5d: %s\n", prefix, ctxLineNum, ctxLine))
		}

		writer.WriteString("```\n\n")
	}
}

func readWordsFromFile(filePath string) ([]string, error) {
	words := []string{}
	file, err := os.Open(filePath)
	if err != nil {
		return words, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return words, err
	}
	return words, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isFileType(path string, types []string) bool {
	ext := filepath.Ext(path)
	for _, t := range types {
		if ext == t {
			return true
		}
	}
	return false
}
