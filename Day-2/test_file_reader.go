package main

import (
	"fmt"
	"os"
	"strings"
)

// testReadFileContent demonstrates the ReadFileContent function with various test cases
func testReadFileContent() {
	fmt.Println("=== ReadFileContent Function Testing ===")

	// Test cases covering all requirements
	testCases := []struct {
		name        string
		filePath    string
		description string
	}{
		{
			name:        "Valid file with content",
			filePath:    "sample.txt",
			description: "Should successfully read the entire file content",
		},
		{
			name:        "Empty file",
			filePath:    "empty.txt",
			description: "Should return empty string without error",
		},
		{
			name:        "Non-existent file",
			filePath:    "nonexistent.txt",
			description: "Should return error: file does not exist",
		},
		{
			name:        "Directory instead of file",
			filePath:    ".",
			description: "Should return error: path is a directory",
		},
		{
			name:        "Empty file path",
			filePath:    "",
			description: "Should return error: file path cannot be empty",
		},
		{
			name:        "File with spaces in path",
			filePath:    "   ",
			description: "Should return error: file path cannot be empty (trimmed)",
		},
	}

	// Run each test case
	for i, tc := range testCases {
		fmt.Printf("Test %d: %s\n", i+1, tc.name)
		fmt.Printf("File Path: %q\n", tc.filePath)
		fmt.Printf("Expected: %s\n", tc.description)

		// Call the ReadFileContent function
		content, err := ReadFileContent(tc.filePath)

		if err != nil {
			fmt.Printf("Result: ERROR - %v\n", err)
		} else {
			// Display content (truncate if too long for readability)
			displayContent := content
			if len(content) > 200 {
				displayContent = content[:200] + "... (truncated)"
			}
			// Replace newlines with \n for better display
			displayContent = fmt.Sprintf("%q", displayContent)
			fmt.Printf("Result: SUCCESS - Content: %s\n", displayContent)
			fmt.Printf("Content Length: %d characters\n", len(content))
		}

		fmt.Println(strings.Repeat("-", 60))
	}

	// Demonstrate the alternative ioutil implementation
	fmt.Println("\n=== Alternative Implementation Test (ReadFileContentIoutil) ===")
	content, err := ReadFileContentIoutil("sample.txt")
	if err != nil {
		fmt.Printf("ReadFileContentIoutil Error: %v\n", err)
	} else {
		fmt.Printf("ReadFileContentIoutil Success: %d characters read\n", len(content))
	}

	// Demonstrate streaming for large files
	fmt.Println("\n=== Streaming File Processing Example ===")
	fmt.Println("Processing sample.txt line by line:")

	lineProcessor := func(line string, lineNumber int) error {
		fmt.Printf("  Line %d: %q\n", lineNumber, line)
		// Process only first 10 lines to avoid spam
		if lineNumber >= 10 {
			fmt.Println("  ... (stopping at line 10 for demo)")
			return fmt.Errorf("demo limit reached") // This will stop the processing
		}
		return nil
	}

	if err := StreamFileLines("sample.txt", lineProcessor); err != nil {
		// We expect an error from our demo limit
		if err.Error() != "error processing line 10: demo limit reached" {
			fmt.Printf("Unexpected streaming error: %v\n", err)
		}
	}

	// Test streaming with empty file
	fmt.Println("\nTesting streaming with empty file:")
	emptyProcessor := func(line string, lineNumber int) error {
		fmt.Printf("  Line %d: %q\n", lineNumber, line)
		return nil
	}

	if err := StreamFileLines("empty.txt", emptyProcessor); err != nil {
		fmt.Printf("Empty file streaming error: %v\n", err)
	} else {
		fmt.Println("  Empty file processed successfully (no lines)")
	}
}

// createTestFiles creates additional test files for comprehensive testing
func createTestFiles() {
	// Create a large test file to demonstrate memory concerns
	largeFile, err := os.Create("large_test.txt")
	if err != nil {
		fmt.Printf("Warning: Could not create large test file: %v\n", err)
		return
	}
	defer largeFile.Close()

	// Write 1000 lines to demonstrate large file handling
	for i := 1; i <= 1000; i++ {
		fmt.Fprintf(largeFile, "This is line %d of the large test file. It contains some content to make it larger.\n", i)
	}

	fmt.Println("Created large_test.txt with 1000 lines for testing")
}

// cleanupTestFiles removes created test files
func cleanupTestFiles() {
	testFiles := []string{"large_test.txt"}

	for _, file := range testFiles {
		if err := os.Remove(file); err != nil {
			// Ignore errors if file doesn't exist
			if !os.IsNotExist(err) {
				fmt.Printf("Warning: Could not remove %s: %v\n", file, err)
			}
		}
	}
}
