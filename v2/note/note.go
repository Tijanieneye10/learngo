package note

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid title or content")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note title %v has the following content:\n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json2.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
