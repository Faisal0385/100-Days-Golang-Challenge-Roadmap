package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// rand/v2 automatically seeds with crypto/rand for you
	target := rand.IntN(100) + 1
	var guess int
	attempts := 0

	fmt.Println("🎯 Guess the Number Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < target {
			fmt.Println("📉 Too low! Try a higher number.")
		} else if guess > target {
			fmt.Println("📈 Too high! Try a lower number.")
		} else {
			fmt.Printf("🎉 Correct! The number was %d. You guessed it in %d attempts.\n", target, attempts)
			break
		}
	}
}
