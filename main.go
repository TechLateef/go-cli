package main

import (
	"fmt"
	"go-cli-learning/cmd"
	"os"
	"strings"
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
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a note to add")
			return
		}
		note := strings.Join(os.Args[2:], " ") // join all the words after "add"
		err := cmd.AddNote(note)
		if err != nil {
			fmt.Println("Error adding note:", err)
		} else {
			fmt.Println("Note added successfully")
		}
	case "list":
		notes, err := cmd.ListNotes()
		if err != nil {
			fmt.Println("Error reading notes:", err)
			return
		}
		fmt.Println("Your Notes:")
		fmt.Println(notes)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])

	}
}
