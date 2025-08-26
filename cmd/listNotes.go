package cmd

import "os"

const notesFile = "notes.txt"

func ListNotes() (string, error) {
	data, err := os.ReadFile(notesFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
