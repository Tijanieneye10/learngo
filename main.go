package main

import "fmt"

// variables

func main() {
	printSomething(40)
	anyParamAndReturn(2)
	anyParamAndReturn("Hello world")
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println(intVal + 1)
	}
}

func anyParamAndReturn(value interface{}) interface{} {
	intVal, isInt := value.(int)
	if isInt {
		fmt.Println(intVal)
	}

	stringVal, isString := value.(string)

	if isString {
		fmt.Println(stringVal)
	}

	return nil
}
