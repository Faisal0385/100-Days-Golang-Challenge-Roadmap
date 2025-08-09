// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}

	fmt.Print("Enter max multiplier (default 10): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	max := 10
	if input != "" {
		if m, e := strconv.Atoi(input); e == nil && m > 0 {
			max = m
		}
	}

	// Print nicely aligned table
	fmt.Printf("\nMultiplication table for %d (1 to %d):\n", n, max)
	for i := 1; i <= max; i++ {
		// use %2d or %3d depending on width you want
		fmt.Printf("%2d x %2d = %3d\n", n, i, n*i)
	}
}
