package main

import "fmt"

func main() {

	accountBalance := 1000.0
	fmt.Println("Welcome to Brainy Bank")

	for {

		fmt.Println("1. Check my account balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Make your choice ")

		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:

			fmt.Printf("Your account balance is N%v", accountBalance)
			continue

		case 2:

			depositAmount := 0.0

			fmt.Println("How much to deposit? ")

			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount stated")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your account balance is now", accountBalance)

		case 3:

			fmt.Println("How much to withdraw? ")
			withdrawalAmount := 0.0
			fmt.Scan(&withdrawalAmount)

			accountBalance -= withdrawalAmount
			fmt.Println("Your new balance is", accountBalance)
			continue
		default:
			fmt.Println("Thanks for banking with us!")
			return

		}
	}

}
