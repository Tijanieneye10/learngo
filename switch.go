package main

import "fmt"

// variables

func main(){

	x := "Male"

	switch x {
	case "Male":
		fmt.Println("The gender of the", x)
		fallthrough // to always run other case.
	case "Female":
		fmt.Println("The gender of the", x)
	default:
		fmt.Println("No gender specified")
	}
}

