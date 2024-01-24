package main

import "fmt"

func main() {
	result := closure(5)

	fmt.Println(result(2))
}

func closure(factory int) func(int) int {
	return func(number int) int {
		return number * factory
	}
}
