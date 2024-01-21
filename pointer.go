package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println(*agePointer)

	checkAge(&age)

}

func checkAge(age *int) {
	fmt.Println("Check age is", *age)
	*age = *age - 4
}
