package cmd

import (
	"fmt"
	"time"
)

const notesFile = "notes.json"

func ListNotes() error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}
	for _, note := range notes {
		fmt.Printf("%d. %s (%s)\n", note.ID, note.Text, note.CreatedAt.Format(time.RFC822))
	}
	return nil
}
