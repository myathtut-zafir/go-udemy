package main

import (
	"bufio"
	"fmt"
	"github.com/pointers/note"
	"github.com/pointers/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
	DisplayNote()
}

func main() {
	title := getUserInput("Enter your title: ")
	desc := getUserInput("Enter your desc: ")
	todoText := getUserInput("Enter your todo: ")

	todo := todo.New(todoText)
	todo.DisplayNote()
	// todo.Save()
	saveData(todo)

	note := note.New(title, desc)
	note.DisplayNote()
	// note.Save()
	saveData(note)
}
func saveData(data saver) {
	data.Save()
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}
