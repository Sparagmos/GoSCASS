# General Indicators

- **TODO**: Marks sections of code that need to be completed or revisited.
- **FIXME**: Indicates code that is broken or needs fixing.
- **BUG**: Highlights known bugs or issues in the code.
- **HACK**: Notes a workaround or temporary solution that may not be ideal.
- **XXX**: Draws attention to a problematic or risky part of the code.
- **DEBUG**: May indicate that debugging code is still present and needs removal.
- **DEPRECATED**: Marks usage of deprecated functions or methods that should be updated.
- **NOTE**: General comments that might contain important information.
- **OPTIMIZE**: Indicates code that can be improved for better performance.

# Security-Related Terms

- **PASSWORD**: Could reveal hardcoded passwords or sensitive information.
- **SECRET**: Similar to PASSWORD, may indicate hardcoded secrets.
- **API_KEY** or **KEY**: Could point to API keys or cryptographic keys embedded in the code.
- **TOKEN**: May reveal authentication tokens hardcoded in the code.
- **PRIVATE_KEY** / **PUBLIC_KEY**: Indicates cryptographic keys that should be securely stored.
- **CREDENTIALS**: May expose user credentials within the code.

# Suspicious Functions and Methods

## General

- **eval**: Execution of code represented as a string can lead to code injection vulnerabilities.
- **exec**: Similar to `eval`, executes commands which may be dangerous if misused.
- **system**: Executes system commands; can be risky if user input is involved.

## Python

- **pickle.loads**: Deserializing data from untrusted sources can execute arbitrary code.
- **os.system**: Executes shell commands from within Python.
- **subprocess.Popen**: Spawns new processes; needs careful input validation.
- **input** (Python 2.x): Can execute user input as code.
- **paramiko**: Used for SSH connections; ensure it's used securely.

## JavaScript

- **eval**: Executes JavaScript code represented as a string.
- **document.write**: Can lead to Cross-Site Scripting (XSS) vulnerabilities.
- **innerHTML**: Direct assignment can introduce XSS if not properly sanitized.

## PHP

- **eval**: Executes PHP code from a string.
- **exec**, **shell_exec**, **passthru**, **system**: Executes shell commands.
- **preg_replace** with `/e` modifier: Allows code execution within regex.

## C/C++

- **gets**: Unsafe function that can lead to buffer overflows.
- **strcpy**, **strcat**: Should be replaced with safer alternatives like `strncpy`, `strncat`.
- **sprintf**: Can cause buffer overflows; consider using `snprintf`.

# Hardcoded Values

- **Hardcoded IPs**: IP addresses directly in the code.
- **Hardcoded URLs**: URLs that may need to be configurable.
- **Connection Strings**: Database connection details hardcoded.

# Logging and Debugging

- **console.log**: Debugging statements in JavaScript that may need removal.
- **print** statements: General debugging outputs in various languages.
- **alert**: JavaScript alerts used for debugging.

# Potential Vulnerabilities

- **SELECT \***: Using `SELECT *` in SQL queries can be inefficient and expose unnecessary data.
- **DROP TABLE**: Dangerous SQL statements that can delete data.
- **No input validation**: Comments indicating lack of validation.
- **race condition**: Indicates potential threading issues.
- **concurrency**: Areas that may need attention for concurrent execution problems.

# Suppressed Warnings

- **#nosec**: Comments used to suppress security warnings; these should be reviewed.
- **@SuppressWarnings**: In Java, may hide important warnings.

# Licensing and Legal

- **LICENSE**: To check for proper license headers.
- **COPYRIGHT**: Verify correct copyright notices.

# Other Indicators

- **Not Implemented**: Indicates incomplete functionality.
- **TBD**: "To Be Determined" sections that need attention.
- **Temporary**: Code that is meant to be temporary but may have been left in.

