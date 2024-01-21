package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"v2/note"
)

func main() {
	title, content := getNoteData()

	noteDetails, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	noteDetails.Display()
	err = noteDetails.Save()

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
