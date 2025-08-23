package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// os.Args holds the cli arguments
	// os.Args[0] is the program name
	// os.Args[1...] = user inputs

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}

	command := args[1] // the first arguemnt after the program name

	switch command {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)

		name := greetCmd.String("name", "World", "a name to greet")
		times := greetCmd.Int("times", 1, "number of times to greet")

		// Parse the flags for "greet"
		greetCmd.Parse(args[2:])
		for i := 0; i < *times; i++ {
			fmt.Printf("Hello, %s! 🎉\n", *name)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)

	}
}
