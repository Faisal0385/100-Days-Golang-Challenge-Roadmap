package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read input from the user (stdin)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a sentence: ")

	// Read user input line
	if scanner.Scan() {
		input := scanner.Text()

		// Split the input into words using strings.Fields
		words := strings.Fields(input)

		// Count the number of words
		wordCount := len(words)

		fmt.Printf("Number of words: %d\n", wordCount)
	} else {
		fmt.Println("No input detected.")
	}
}
