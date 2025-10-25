package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/pointers/note"
	"github.com/pointers/todo"
)

type saver interface {
	Save() error
}

type outputter interface {
	saver
	DisplayNote()
}

func main() {
	printSomethings("hi")
	title := getUserInput("Enter your title: ")
	desc := getUserInput("Enter your desc: ")
	todoText := getUserInput("Enter your todo: ")

	todo := todo.New(todoText)
	// todo.DisplayNote()
	// todo.Save()
	saveData(todo)
	output(todo)

	note := note.New(title, desc)
	// note.DisplayNote()
	// note.Save()
	saveData(note)
	output(note)
}
func saveData(data outputter) {
	data.Save()
}
func output(data outputter) {
	data.DisplayNote()
}

func printSomethings(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		intVal++
		fmt.Println("This is an integer:", intVal)
		return
	}
	// switch value.(type){
	// 	case string:
	// 		fmt.Println("This is a string")
	// 	case int:
	// 		fmt.Println("This is an integer")
	// 	case bool:
	// 		fmt.Println("This is a boolean")
	// 	default:
	// 		fmt.Println("Unknown type")
	// }
	fmt.Println(value)
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
