package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a positive number")
		return
	}

	// Create an array/slice to hold the sequence
	fib := make([]int, n)

	// First two numbers
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}

	// Loop to generate sequence
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	// Output the sequence
	fmt.Println("Fibonacci Sequence:", fib)
}
