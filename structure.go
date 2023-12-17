package main

import "fmt"

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}


func main() {
	tijani := Passenger{
		Name: "Usman Tijani",
		TicketNumber: 1,
		Boarded: false,
	}
	fmt.Println(tijani)

	var (
		eneye = Passenger{Name: "Eneye", TicketNumber: 2, Boarded: true}
		halimat = Passenger{Name: "Halimiat", TicketNumber: 3, Boarded: false}
	)

	fmt.Println(eneye, halimat)

	var hassan Passenger

	hassan.Name = "Hassan"
	hassan.TicketNumber = 5
	hassan.Boarded = true

	if(hassan.Boarded){
		fmt.Println("Yeah Hassan has already boarded the bus")
	}
}