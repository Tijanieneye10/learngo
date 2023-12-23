package main

import "fmt"


func main(){
	name := "Tijani"

	var copyIt *string = &name

	fmt.Println(copyIt)
}