package main

import "fmt"

func playFn() {
	add, multiply := multiplyAndAddition(5, 10)

	fmt.Println(add, multiply)

	var result = 200

	if result > 100 {
		fmt.Println("Good result")
	} else {
		fmt.Println("Bad result")
	}
}

func multiplyAndAddition(first, second int) (addition, multiply int) {
	addition = first + second
	multiply = first * second
	return
}
