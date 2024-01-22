package main

import "fmt"

func main() {
	prices := []float64{10.99, 2.1, 5.4, 9.3, 42.5}

	fmt.Println(prices[1:2])

	//Start from beginning and end at the third
	fmt.Println(prices[:3])

	//Start from 2 to the last element
	fmt.Println(prices[2:])
}
