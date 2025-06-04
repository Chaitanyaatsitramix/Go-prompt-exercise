package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFileContent reads the entire content of a file specified by filePath
// and returns it as a string along with any error that might occur.
//
// Parameters:
//   - filePath: string - The path to the file to be read (relative or absolute)
//
// Returns:
//   - string: The complete content of the file as a string
//   - error: Any error that occurred during the file reading process, or nil if successful
//
// This function demonstrates several important Go concepts:
// 1. File handling with proper resource management
// 2. Error handling patterns
// 3. The defer statement for cleanup
// 4. Multiple return values (idiomatic Go pattern)
func ReadFileContent(filePath string) (string, error) {
	// Input validation: Check if the file path is empty
	// Why check this? An empty path would cause os.Open to fail with a confusing error.
	// It's better to catch this early and provide a clear error message.
	if filePath == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}

	// Step 1: Open the file
	// os.Open() opens the file for reading and returns a *File and an error
	// Why use os.Open instead of os.OpenFile? os.Open is simpler for read-only operations
	// and is equivalent to os.OpenFile(name, O_RDONLY, 0)
	file, err := os.Open(filePath)
	if err != nil {
		// If os.Open fails, it could be due to:
		// - File doesn't exist (os.ErrNotExist)
		// - Permission denied
		// - Path is a directory, not a file
		// - Invalid path format
		// We wrap the error to provide context about what we were trying to do
		return "", fmt.Errorf("failed to open file '%s': %w", filePath, err)
	}

	// Step 2: Ensure the file is closed when the function returns
	// defer schedules file.Close() to be called when ReadFileContent returns
	// Why use defer?
	// 1. Ensures cleanup happens even if an error occurs later
	// 2. Keeps the cleanup code close to the resource acquisition
	// 3. Prevents resource leaks (file handles are limited system resources)
	// 4. Makes the code more readable and maintainable
	defer func() {
		// We call Close() and handle its potential error
		// While reading files rarely fails to close, it's good practice to check
		if closeErr := file.Close(); closeErr != nil {
			// In a more sophisticated application, you might log this error
			// For this example, we'll just print it
			fmt.Printf("Warning: failed to close file '%s': %v\n", filePath, closeErr)
		}
	}()

	// Step 3: Get file information to check if it's actually a file
	// Why do this? os.Open can successfully open directories on some systems
	// but we want to read file content, not directory listings
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info for '%s': %w", filePath, err)
	}

	// Check if the path points to a directory instead of a regular file
	if fileInfo.IsDir() {
		return "", fmt.Errorf("'%s' is a directory, not a file", filePath)
	}

	// Step 4: Read the file content using bufio.Scanner
	// Why use bufio.Scanner instead of ioutil.ReadAll or io.ReadAll?
	// 1. Scanner provides line-by-line reading with built-in buffering
	// 2. More memory efficient for large files
	// 3. Handles different line endings automatically
	// 4. Provides more control over the reading process

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// We'll build the content string by reading line by line
	var content string
	lineCount := 0

	// Read the file line by line
	// Why loop instead of reading all at once?
	// This approach is more memory efficient and allows us to process very large files
	for scanner.Scan() {
		lineCount++
		// scanner.Text() returns the current line without the line ending
		line := scanner.Text()

		// Add the line to our content string
		// For the first line, don't add a newline prefix
		if lineCount == 1 {
			content = line
		} else {
			// Add newline + line content
			// We reconstruct newlines because scanner.Text() strips them
			content += "\n" + line
		}
	}

	// Step 5: Check if the scanner encountered any errors during reading
	// Why check scanner.Err()? The scanner might fail due to:
	// - I/O errors during reading
	// - Encoding issues (if file contains invalid UTF-8)
	// - Other system-level problems
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file '%s': %w", filePath, err)
	}

	// If we reach here, the file was read successfully
	// Return the content and nil error (indicating success)
	return content, nil
}

// Alternative implementation using os and io packages instead of bufio
// This demonstrates a different approach to the same problem
func ReadFileContentSimple(filePath string) (string, error) {
	// Input validation
	if filePath == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}

	// Read the entire file at once using os.ReadFile (Go 1.16+)
	// This is simpler but less memory efficient for large files
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	// Convert bytes to string and return
	return string(data), nil
}

// ReadFileContentIoutil demonstrates an alternative implementation using ioutil
// This is kept for educational purposes and comparison
func ReadFileContentIoutil(filePath string) (string, error) {
	// Validate input
	if strings.TrimSpace(filePath) == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}

	// Check if file exists and is not a directory
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %s", filePath)
		}
		return "", fmt.Errorf("failed to access file info for %s: %w", filePath, err)
	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("path is a directory, not a file: %s", filePath)
	}

	// For empty files, return empty string
	if fileInfo.Size() == 0 {
		return "", nil
	}

	// Using os.Open with manual reading (alternative to ioutil.ReadFile)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsPermission(err) {
			return "", fmt.Errorf("permission denied: cannot read file %s", filePath)
		}
		return "", fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Read all content at once
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return string(content), nil
}

// How would you refactor this for streaming large files line-by-line?
// StreamFileLines demonstrates how to handle large files by processing them line by line
// instead of loading everything into memory at once
func StreamFileLines(filePath string, processor func(line string, lineNumber int) error) error {
	// Validate input
	if strings.TrimSpace(filePath) == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// Check file existence and type
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filePath)
		}
		return fmt.Errorf("failed to access file info for %s: %w", filePath, err)
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("path is a directory, not a file: %s", filePath)
	}

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied: cannot read file %s", filePath)
		}
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Process file line by line using bufio.Scanner
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		// Call the processor function for each line
		if err := processor(line, lineNumber); err != nil {
			// If processor returns an error, we can choose to stop or continue
			// In this case, we stop and return the error
			return fmt.Errorf("error processing line %d: %w", lineNumber, err)
		}

		lineNumber++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return nil
}

/*
COMPREHENSIVE ANSWERS TO YOUR QUESTIONS:

Q: How does defer work and why is it important here?
A: defer schedules a function call to be executed when the surrounding function returns,
   regardless of how it returns (normal return, panic, etc.). It's important here because:
   1. It ensures the file is always closed, preventing resource leaks
   2. It keeps cleanup code close to resource allocation (better readability)
   3. It handles cleanup even if errors occur later in the function
   4. It follows Go's idiom of "acquire resource, defer cleanup, use resource"

Q: What happens if the file doesn't exist?
A: os.Open() will return an error (typically os.ErrNotExist wrapped in a *PathError).
   We check for this error and return early with a descriptive error message.
   The defer statement ensures any cleanup is still performed if needed.

Q: How do we handle different error cases?
A: We handle several error cases:
   1. Empty file path - early validation with clear error message
   2. File doesn't exist - os.Open returns an error, we wrap it with context
   3. Permission denied - also caught by os.Open, wrapped with context
   4. Path is directory - we check fileInfo.IsDir() and return specific error
   5. Reading errors - scanner.Err() catches I/O errors during reading
   6. Close errors - handled in defer with warning (non-fatal)

Q: What's the difference between different reading approaches?
A: Three main approaches demonstrated:
   1. bufio.Scanner (used in main function):
      - Line-by-line reading with automatic line ending handling
      - Memory efficient for large files
      - More control over reading process
      - Good for processing structured text

   2. os.ReadFile (shown in alternative):
      - Reads entire file at once
      - Simpler code, less control
      - Uses more memory for large files
      - Good for small to medium files

   3. io.ReadAll (not shown but mentioned):
      - Similar to os.ReadFile but works with any io.Reader
      - Good when you already have an open file handle

Q: When should you use each approach?
A:
   - Use bufio.Scanner when: processing large files, need line-by-line processing,
     want memory efficiency, or need to handle special line endings
   - Use os.ReadFile when: files are small-medium size, you need simplicity,
     and memory usage isn't a concern
   - Use io.ReadAll when: you already have an io.Reader and want to read it completely

EXAMPLE USAGE:
	content, err := ReadFileContent("example.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", content)
*/
