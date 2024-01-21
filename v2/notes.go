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

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)
	//var value string
	//fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
