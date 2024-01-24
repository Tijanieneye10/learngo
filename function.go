package main

import "fmt"

type doubleFn func(int) int

// variables

func main() {

	fmt.Println(isWorking(4, 6))

	a, b, c := mulltipleReturn(6, 7)

	fmt.Println(a, b, c)

	numbers := []int{4, 5, 7, 8}

	fmt.Println(doubleSlice(&numbers, double))

}

func isWorking(first, second int) int {
	return first + second
}

// multiple returns
func mulltipleReturn(first, second int) (int, int, string) {
	return first, second, "testing"
}

//Ignore values using _

// Function as a value
func doubleSlice(numbers *[]int, double doubleFn) []int {
	var values []int

	for _, el := range *numbers {
		values = append(values, double(el))
	}

	return values
}

//Another way to do it.
//func doubleSlice(numbers *[]int, double func(int) int) []int {
//	var values []int
//
//	for _, el := range *numbers {
//		values = append(values, double(el))
//	}
//
//	return values
//}

func double(number int) int {
	return number * 2
}

// **** Func return another func
func returnFnAsValue() doubleFn {
	return double
}
