package main

import (
	"errors"
	"fmt"
)

type str string

type UserDetail struct {
	Email    str
	Password string
	OtherDetail
}

type OtherDetail struct {
	Age     int
	Address string
	balance float64
}

func (user UserDetail) login() (string, error) {
	if user.Email == "john@gmail.com" && user.Password == "12345" {
		return "User login successfully", nil
	}

	return "", errors.New("invalid email or password")
}

func main() {
	user := UserDetail{
		Email:    "john@gmail.com",
		Password: "12345",
		OtherDetail: OtherDetail{
			Age:     25,
			Address: "No 2 John Doe House",
		},
	}

	fmt.Println(user.Address, user.Email)

	login, err := user.login()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(login)
}
