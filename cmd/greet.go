package cmd

import (
	"flag"
	"fmt"
)

// Greet print a greeting message with a name flag

func Greet(args []string) {
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	name := greetCmd.String("name", "World", "a name to greet")
	times := greetCmd.Int("times", 1, "number of time to greet")

	// Parse only the args for greet
	greetCmd.Parse(args)

	for i := 0; i < *times; i++ {
		fmt.Printf("Hello, %s! 🎉\n", *name)
	}
}
