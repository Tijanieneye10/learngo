package main


import ("fmt")

func printSlices(title string, slices []string){
	fmt.Println(title)

	for i := 0; i < len(slices); i++ {
		element := slices[i]
		fmt.Println(element)
	}

}


func main(){
	route := []string{"home", "about", "contact"}

	route = append(route, "services")

	printSlices("Route is my Routes", route);

	route = route[1:3]

	printSlices("Second routes", route);
}