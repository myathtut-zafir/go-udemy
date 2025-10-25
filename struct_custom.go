package main

import (
	"bufio"
	"fmt"
	"github.com/pointers/note"
	"os"
	"strings"
)

func main() {
	title := getUserInput("Enter your title: ")
	desc := getUserInput("Enter your desc: ")

	note := note.New(title, desc)
	note.DisplayNote()
	note.Save()
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ",prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}
