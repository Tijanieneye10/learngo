package main

import "fmt"

type Product struct {
	name  string
	price int
}

func calculateProduct(list [4]Product) {

	var totalCost, totalItem int

	for i := 0; i < len(list); i++ {

		item := list[i]

		totalCost += item.price

		if item.name != " " {
			totalItem += 1
		}
	}

	fmt.Println("Total Cost: ", totalCost)
	fmt.Println("Total Item: ", totalItem)
}

func main() {
	items := [4]Product{
		{name: "Laptop", price: 200000},
		{name: "Bag", price: 200},
		{name: "Phone", price: 310},
	}

	calculateProduct(items)

	// Array
	names := [3]string{"John", "Doe", "Taylor"}
	fmt.Println(names)

	//Slice
	anotherNames := []string{"John", "Doe", "Taylor"}
	fmt.Println(anotherNames)

	// Make Function
	secondNames := make([]string, 3, 5)

	secondNames = []string{"John", "Doe", "Taylor"}

	fmt.Println(secondNames)

}
