package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a word or sentence: ")
	// Read full line (including spaces)
	fmt.Scanln(&input)

	// Normalize: lowercase + remove spaces
	normalized := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	// Check palindrome
	if isPalindrome(normalized) {
		fmt.Println("✅ It's a palindrome!")
	} else {
		fmt.Println("❌ Not a palindrome.")
	}
}

func isPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
