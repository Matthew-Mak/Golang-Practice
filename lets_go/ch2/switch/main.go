package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Red"

	switch language {
	case "Red":
		fmt.Println("Your limit is first layer")
		// break - is added automatically
	case "Blue", "Blue Whistle":
		fmt.Println("Your limit is beyond the start level but proceed with caution!")
	default:
		fmt.Println("Other whistles will help you on your way!")
	}

	n := 5
	switch true { // bool switch
	case n%2 == 0:
		fmt.Println("Even num")
	case n%2 != 0:
		fmt.Println("Odd num")
	default:
		fmt.Println("Never reached")
	}

	hour := time.Now().Hour()
	//fmt.Println(hour)
	//hour := 15

	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
