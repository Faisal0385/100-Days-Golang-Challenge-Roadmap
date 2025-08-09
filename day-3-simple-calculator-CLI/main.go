package main

import (
	"flag"
	"fmt"
	"os"
)

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {
	// Define CLI flags
	num1 := flag.Float64("num1", 0, "First number")
	num2 := flag.Float64("num2", 0, "Second number")
	op := flag.String("op", "+", "Operator (+, -, *, /)")

	// Parse CLI flags
	flag.Parse()

	// Perform calculation
	result, err := calculate(*num1, *num2, *op)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)
}
