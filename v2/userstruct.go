package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthday  time.Time
}

func main() {
	params, err := newUser(
		"Tijani Eneye",
		"Usman",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	params.learnstruct()
	fmt.Println(params.firstname)
	params.clearUsername()
	fmt.Println(params.firstname)
}

// constructor
func newUser(firstname, lastname string) (*User, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("params are required")
	}
	return &User{
		firstname: firstname,
		lastname:  lastname,
	}, nil
}

func (firstUser User) learnstruct() {
	fmt.Println(firstUser.firstname, firstUser.lastname, firstUser.birthday)
}

func (firstUser *User) clearUsername() {
	firstUser.firstname = ""
}
