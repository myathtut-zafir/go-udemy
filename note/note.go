package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func New(title, description string) Note {
	return Note{
		Title:       title,
		Description: description,
	}
}

func (note Note) DisplayNote() {
	fmt.Println("Title:", note.Title)
	fmt.Println("Description:", note.Description)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
