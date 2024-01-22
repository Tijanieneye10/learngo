package main

import "fmt"

func main() {

	hobbies := [3]string{"Coding", "Learning", "Playing"}

	firstEl := hobbies[0]

	secondThirdEl := hobbies[1:3]

	firstSecondEl := hobbies[0:2]

	fmt.Println(hobbies, firstEl, secondThirdEl, firstSecondEl)

	products := []Product{
		{
			Id:    1,
			Name:  "Laptop",
			Price: 1200.99,
		},
		{
			Id:    2,
			Name:  "Monitor",
			Price: 120.05,
		},
	}

	secondProducts := []Product{
		{
			Id:    4,
			Name:  "House",
			Price: 12900.99,
		},
		{
			Id:    5,
			Name:  "Television",
			Price: 121.05,
		},
	}

	anotherProduct := Product{
		Id:    3,
		Name:  "Fan",
		Price: 9.99,
	}

	products = append(products, anotherProduct)

	products = append(products, secondProducts...)

	fmt.Println(products)

	for el, product := range products {
		fmt.Println(el, product.Name)
	}
}

type Product struct {
	Id    int
	Name  string
	Price float64
}
