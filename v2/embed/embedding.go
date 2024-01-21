package main

import "time"

type User struct {
	Firstname string
	Lastname string
	Dob time.Time
}

type Admin struct {
	Email string
	Password string
	User
}

func NewAdmin(email, password) Admin {
	Email: email,
	Password: password,
		User: User{
		Firstname: "John",
		Lastname: "Doe",
		Dob: time.Now()
	},
}

func main() {

}
