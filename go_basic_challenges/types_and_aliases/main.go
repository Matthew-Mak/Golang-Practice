package main

import "fmt"

type duration int

type mile float64
type kilometer float64

const m2km = 1.609

func main() {
	// type duration int
	// var hour duration
	// println(hour)
	// fmt.Printf("%T\n", hour)
	// hour = 3600
	// println(hour)

	// var hour duration = 3600
	// minute := 60
	// fmt.Println(hour != duration(minute))

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Println(kmBerlinToParis)
}
