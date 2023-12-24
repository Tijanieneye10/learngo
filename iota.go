package main

import "fmt"

func main() {

	// Iota is a way of incrementing constants.
	// Short iota
	const (
		Online = iota
		Offline
		Away
		Coming
	)

	fmt.Println(Online, Offline)

	// Long iota
	const (
		Completed = iota
		InProgress = iota
		Working = iota
	)

	//Skip some
	const (
		Going = iota
		_
		Man
	)
}
