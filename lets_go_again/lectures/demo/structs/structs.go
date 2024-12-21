package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	kaly := Passenger{"Kaly", 1, false}
	fmt.Println(kaly)

	var (
		john  = Passenger{Name: "John", TicketNumber: 2}
		surcy = Passenger{Name: "Surcy", TicketNumber: 3}
	)
	fmt.Println(john, surcy)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	kaly.Boarded = true
	john.Boarded = true
	if john.Boarded {
		fmt.Println("Go abord, John!")
	}
	if kaly.Boarded {
		fmt.Println("Go abord, Kaly!")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "Is in the front seat")
}
