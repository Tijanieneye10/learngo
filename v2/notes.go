package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"v2/note"
	todo2 "v2/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	Save() error
	Display()
}

// Implementing interface
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed...")
		return err
	}

	return nil

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func main() {

	title, content := getNoteData()

	todoText := getUserInput("What is your todo? ")

	todo, err := todo2.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
	}

	noteDetails, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	err = outputData(noteDetails)

	if err != nil {
		fmt.Println("Saving the note failed...")
		return
	}

	fmt.Println("Note saved successfully")

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
