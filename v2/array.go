package main

import "fmt"

func main() {
	prices := []float64{10.99, 2.1, 5.4, 9.3, 42.5}

	fmt.Println(prices[1:2])

	//Start from beginning and end at the third
	fmt.Println(prices[:3])

	//Start from 2 to the last element
	fmt.Println(prices[2:])

	names := []string{"John", "Doe", "Example", "Test"}

	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(names), cap(names))

	numbers := []int{3, 4, 5}
	appendedValue := append(numbers, 6)

	fmt.Println(appendedValue)
}
