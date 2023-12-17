package main

import "fmt"

// variables

func variable(){
	// var example = 3

	// compound variable
	var a, b, c = 4, "man", 4.2

	d, e := 3, "Create variable"

	// block variable
	var (
		man string = "human"
		woman string = "woman"

	)

	const (
		Monday = "Monday"
		Tuesday = "Tuesday"
	)

	fmt.Println(a, b, c, man, woman)
}

