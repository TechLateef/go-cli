package cmd

import (
	"encoding/json"
	"os"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func LoadNotes() ([]Note, error) {
	var notes []Note

	file, err := os.ReadFile(notesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{}, nil // Return empty slice if file doesn't exist
		}
		return nil, err
	}

	err = json.Unmarshal(file, &notes)
	return notes, err

}
