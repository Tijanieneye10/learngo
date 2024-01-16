package main

import (
	"errors"
	"fmt"
)


type Person struct {
	Name string
	Age int
}


func (p *Person) String() (int, error) {
	if p.Age < 20 {
		return 0, errors.New("Age must be greater than 20")
	} else {
		return p.Age, nil
	}
}


func main(){
	person := &Person{
		Name: "Tijani",
		Age: 30,
	}

	result, err := person.String()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}
}