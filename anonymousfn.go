package main

import "fmt"

type numberFn func(int) int

func main() {
	numbers := []int{4, 6, 7, 8, 9}

	result := doubleMyNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(result)
}

func doubleMyNumbers(numbers *[]int, double func(int) int) []int {
	var dNumbers []int
	for _, el := range *numbers {
		dNumbers = append(dNumbers, double(el))
	}
	return dNumbers
}
