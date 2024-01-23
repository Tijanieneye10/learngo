package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	person := make(map[string]string)

	person["Name"] = "John Doe"

	fmt.Println(person)

	names := [2]string{"John Doe", "Goodluck"}

	fmt.Println(names[0])

	links := map[string]string{
		"google": "https://www.google.com",
		"aws":    "https://www.aws.com",
	}

	links["linkedin"] = "https://www.linkedin.com"

	delete(links, "google")

	fmt.Println(links)

	//receive the slice type, the initial length, and capacity
	userNames := make([]string, 3, 4)

	userNames[0] = "Test"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Doe")

	fmt.Println(userNames)

	//using make for map only take initial length
	courses := make(map[string]string, 10)

	fmt.Println(courses)

	//Type alias for map
	ratings := make(floatMap, 3)

	ratings["php"] = 9.0
	ratings["go"] = 7.0
	ratings["vue"] = 5.0

	ratings.output()

}
