package main

import "fmt"

const (
	Add = iota
	Subtract
	Divide
	Multiply
)

type Operator int

func (operator Operator) calculate(first, second int) int {
	switch operator {
	case Add:
		return first + second
	case Subtract:
		return first - second
	case Divide:
		return first / second
	case Multiply:
		return first * second
	}
	panic("un handled error")
}

func main(){
	add := Operator(Add)
	fmt.Println(add.calculate(5, 10))

	subtract := Operator(Subtract)
	fmt.Println(subtract.calculate(22, 10))

}