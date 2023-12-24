package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

func main() {
	num1 := []int{1, 2, 3, 4}
	num2 := []int{5,6,4,1,5}

	all := append(num1, num2...)

	fmt.Println(sum(all...))
}