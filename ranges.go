package main

import "fmt"

func main() {
	names := []string{"foo", "bar", "baz"}

	for _, element := range names {
		fmt.Println(element)

		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}
}
      