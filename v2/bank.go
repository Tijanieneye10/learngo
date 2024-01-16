package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}

func main() {

	fmt.Println("Here is the balance", getBalanceFromFile())

	accountBalance := getBalanceFromFile()
	fmt.Println("Welcome to Brainy Bank")

	for {

		welcomeMessage()

		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:

			fmt.Printf("Your account balance is N%v", accountBalance)
			continue

		case 2:

			depositAmount := 0.0

			fmt.Print("How much to deposit? ")

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount stated")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your account balance is now", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:

			fmt.Print("How much to withdraw? ")
			withdrawalAmount := 0.0
			fmt.Scan(&withdrawalAmount)

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Your new balance is", accountBalance)
			continue
		default:
			fmt.Println("Thanks for banking with us!")
			return

		}
	}

}
