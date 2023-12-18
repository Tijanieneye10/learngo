package main

import "fmt"

func main() {
	// array := []string{"hello", "world", "man"}

	// for i, element := range array {
	// 	fmt.Println(i, element)
	// }

	rooms := [...]Rooms{
		{name: "Warehouse", cleaned: true},
		{name: "Office", cleaned: false},
		{name: "House", cleaned: true},
		{name: "Workshop", cleaned: true},
	}

	rooms[1].cleaned = true

	checkIsClean(rooms)
}

type Rooms struct {
	name    string
	cleaned bool
}

func checkIsClean(rooms [4]Rooms){
	for i := 0; i< len(rooms); i++ {
		room := rooms[i]

		if(room.cleaned) {
			fmt.Println("The room is cleaned")
		} else {
			fmt.Println("The room is not cleaned")
		}
	}
}
