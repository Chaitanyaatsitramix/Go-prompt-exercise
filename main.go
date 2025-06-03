package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Go Prompt Exercise - All Features ===\n")

	// Feature 1: Hello World functionality (from second prompt)
	fmt.Println("1. Hello World Feature:")
	printHelloWorld()

	// Feature 2: Server message (from first prompt)
	fmt.Println("\n2. Server Feature:")
	printServerMessage()

	// Feature 3: Addition function (from current prompt)
	fmt.Println("\n3. Addition Function Feature:")
	demonstrateAddition()
}

// printHelloWorld - functionality from the second prompt
func printHelloWorld() {
	// Step 1: Use fmt.Println to print the string to the console
	// fmt.Println automatically adds a newline character at the end
	fmt.Println("Hello, World!")

	// Step 2: Additional demonstration of string printing
	// This shows an alternative way to print using fmt.Print (without automatic newline)
	fmt.Print("This is printed using fmt.Print ")
	fmt.Print("(no automatic newline)\n")

	// Step 3: Using fmt.Printf for formatted output
	// %s is a placeholder for string values
	message := "Hello, World!"
	fmt.Printf("Using fmt.Printf with variable: %s\n", message)
}

// printServerMessage - functionality from the first prompt
func printServerMessage() {
	fmt.Println("Starting server on :8080...")
	fmt.Println("Server would be running at: http://localhost:8080")
	fmt.Println("Available endpoints would be:")
	fmt.Println("  - GET / (Welcome message)")
	fmt.Println("  - GET /health (Health check)")
	fmt.Println("Note: Server is not actually started in this demo")
}

// demonstrateAddition - functionality from the current prompt
func demonstrateAddition() {
	// Example usage of the addNumbers function
	// Define two integer variables to demonstrate the function
	firstNumber := 15
	secondNumber := 25

	// Call the addNumbers function with the two integers
	// and store the result in a variable
	result := addNumbers(firstNumber, secondNumber)

	// Print the original numbers and their sum
	fmt.Printf("First number: %d\n", firstNumber)
	fmt.Printf("Second number: %d\n", secondNumber)
	fmt.Printf("Sum: %d + %d = %d\n", firstNumber, secondNumber, result)

	// Additional examples with different numbers
	fmt.Println("\nAdditional examples:")
	fmt.Printf("10 + 5 = %d\n", addNumbers(10, 5))
	fmt.Printf("100 + 200 = %d\n", addNumbers(100, 200))
	fmt.Printf("-5 + 15 = %d\n", addNumbers(-5, 15))
	fmt.Printf("0 + 42 = %d\n", addNumbers(0, 42))
}

// addNumbers takes two integer parameters (a and b) and returns their sum
// Parameters:
//   - a: the first integer to add
//   - b: the second integer to add
//
// Returns:
//   - int: the sum of a and b
func addNumbers(a int, b int) int {
	// Addition logic: Use the + operator to add the two integers
	// The + operator performs mathematical addition on the numeric values
	// and returns the result as an integer
	sum := a + b

	// Return the calculated sum to the caller
	return sum
}
