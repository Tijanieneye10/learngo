package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := passNumber(5.2)
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong")
	}

	fmt.Println(result)
}

func passNumber(number float64) (float64, error) {
	if number < 2 {
		return 0, errors.New("number can't be less than 2")
	}

	return number, nil
}
