package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go World!!")
	distance := 60.8 // distance in km
	var kilo int = 90
	fmt.Println(kilo * 2)
	fmt.Printf("The distance in miles is %f \n\n", distance*0.621)
}
