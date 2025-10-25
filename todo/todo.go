package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) Todo {
	return Todo{
		Text: text,
	}
}

func (todo Todo) DisplayNote() {
	fmt.Println("Text:", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
