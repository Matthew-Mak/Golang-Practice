package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("moderately priced")
	default:
		fmt.Println("?")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seat")
	case Business:
		fmt.Println("Business seat")
	case FirstClass:
		fmt.Println("FirstClass seat")
	default:
		fmt.Println("Other seating")
	}

}
