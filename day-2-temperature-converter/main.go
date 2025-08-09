package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp float64
	var unit string

	fmt.Print("Enter temperature (number): ")
	_, err := fmt.Scan(&temp)
	if err != nil {
		fmt.Println("Invalid input for temperature.")
		return
	}

	fmt.Print("Enter unit (C or F): ")
	_, err = fmt.Scan(&unit)
	if err != nil {
		fmt.Println("Invalid input for unit.")
		return
	}

	unit = strings.ToUpper(unit)

	if unit == "C" {

		// Celsius to Fahrenheit
		fahrenheit := (temp * 9 / 5) + 32
		fmt.Printf("%.2f°C is %.2f°F\n", temp, fahrenheit)

	} else if unit == "F" {

		// Fahrenheit to Celsius
		celsius := (temp - 32) * 5 / 9
		fmt.Printf("%.2f°F is %.2f°C\n", temp, celsius)

	} else {
		fmt.Println("Unknown unit. Please enter C or F.")
	}
}
