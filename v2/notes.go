package main

import (
	"fmt"
	"v2/note"
)

func main() {
	title, content := getNoteData()

	noteDetails, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	noteDetails.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note title ")
	content := getUserInput("Note content")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
