package cmd

import (
	"encoding/json"
	"os"
)

func SaveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(notesFile, data, 0644)
}
