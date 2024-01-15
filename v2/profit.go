package main

import "fmt"

func main(){

	var revenue float64
	var expenses float64

	const PERCENTAGE float64 = 100

	fmt.Print("What is your revenue? ")
	fmt.Scan(&revenue)

	fmt.Print("What is your expenses? ")
	fmt.Scan(&expenses)


	profit := revenue - expenses

	profitMargin := (profit/revenue) * PERCENTAGE

	formattedProfitMargin := fmt.Sprintf("%.2f", profitMargin)

	fmt.Println("Your profit margin is", formattedProfitMargin)

	fmt.Println("You make a profit of", profit)
}