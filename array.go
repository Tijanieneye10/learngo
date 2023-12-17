package main

import "fmt"

func main() {
	array := []string {"hello", "world", "man"}

	for i, element := range array {
		fmt.Println(i, element)
	}
}