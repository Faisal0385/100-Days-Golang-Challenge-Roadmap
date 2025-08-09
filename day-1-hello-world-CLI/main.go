package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	lang := flag.String("lang", "en", "Language for the greeting: en, es, bn")
	name := flag.String("name", "", "Name to greet")

	// Parse flags from CLI input
	flag.Parse()

	// Choose greeting based on language
	greeting := "Hello"
	switch *lang {
	case "es":
		greeting = "Hola"
	case "bn":
		greeting = "হ্যালো"
	case "en":
		// default already set
	default:
		greeting = "Hello"
	}

	// Prepare final message
	if *name != "" {
		fmt.Printf("%s, %s!\n", greeting, *name)
	} else {
		fmt.Println("Hello, World!")
	}
}
