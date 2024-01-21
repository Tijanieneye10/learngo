package todo

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"title"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid title or content")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todos.json"
	json, err := json2.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
