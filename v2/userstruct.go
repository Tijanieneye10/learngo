package main

import (
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthday  time.Time
}

func main() {
	params := User{
		firstname: "Tijani Eneye",
		lastname:  "Usman",
		birthday:  time.Now(),
	}

	params.learnstruct()
}

func (firstUser User) learnstruct() {
	fmt.Println(firstUser.firstname, firstUser.lastname, firstUser.birthday)
}
