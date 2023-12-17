package main

import "fmt"

// variables

func main(){
	fmt.Println(isWorking(4,6))

	a,b,c := mulltipleReturn(6,7)

	fmt.Println(a,b,c)
}

func isWorking(first, second int) int {
	return first + second
}

// multiple returns
func mulltipleReturn(first, second int) (int, int, string) {
	return first, second, "testing"
}

//Ignore values using _