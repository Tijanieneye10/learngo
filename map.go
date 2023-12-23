package main

import "fmt"


func main(){
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["bread"] = 10

	shoppingList["eggs"] += 1

	delete(shoppingList, "bread")

	fmt.Println(shoppingList)
}