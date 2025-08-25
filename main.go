package main

import (
	"fmt"
	"go-cli-learning/cmd"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("User: go run main.go <command> [options]")
		return
	}

	switch os.Args[1] {
	case "hello":
		cmd.Hello()
	case "greet":
		cmd.Greet(os.Args[2:])
	case "goodbye":
		cmd.Goodbye(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])

	}
}
