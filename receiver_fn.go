package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func (coord *Coordinate) shift(x, y int){
	coord.X += x
	coord.Y += y

	fmt.Println(coord.X)
	fmt.Println(coord.Y)
}

func main(){
	coord := &Coordinate{5,5}
	coord.shift(1,1)

	fmt.Println(coord.X)
}