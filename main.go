package main

import (
	"fmt"
	"go-cli-learning/cmd"
	"os"
	"strconv"
	"strings"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}
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
		fmt.Println("Your Notes:")
		err := cmd.ListNotes()
		if err != nil {
			fmt.Println("Error reading notes:", err)
			return
		}
	case "timer":
		if len(os.Args) < 3 {
			fmt.Println("Please probide seconds: cli timer 5")
			return
		}
		secs, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid number of seconds:", os.Args[2])
			return
		}
		done := make(chan bool)
		go cmd.StartTimer(secs, done) // run timer in a goroutine

		// Keep main alive until time finishes
		<-done
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])

	}
}
