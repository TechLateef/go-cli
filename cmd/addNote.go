package cmd

import (
	"time"
)

func AddNote(text string) error {
	notes, err := LoadNotes()
	if err != nil {
		return err
	}
	newNote := Note{
		ID:        len(text) + 1,
		Text:      text,
		CreatedAt: time.Now(),
	}
	notes = append(notes, newNote)
	return SaveNotes(notes)
}
