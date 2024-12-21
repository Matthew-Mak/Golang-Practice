//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Addresser interface {
	Address()
}

type Motorcycle struct {
	modelName string
}

type Car struct {
	modelName string
}

type Truck struct {
	modelName string
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m.modelName))
}

func (m Car) String() string {
	return fmt.Sprintf("Car: %v", string(m.modelName))
}

func (m Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(m.modelName))
}

func (m Motorcycle) Address() {
	fmt.Printf("%v, Please proceed to the lift for small vehicles\n", m)
}

func (m Car) Address() {
	fmt.Printf("%v, Please proceed to the lift for standard vehicles\n", m)
}

func (m Truck) Address() {
	fmt.Printf("%v, Please proceed to the lift for large vehicles\n", m)
}

func main() {
	truck := Truck{modelName: "Rim Biliton"}
	truck.Address()
	car := Car{modelName: "Suzuki Gimini"}
	car.Address()
	moto := Motorcycle{modelName: "Kavasaki"}
	moto.Address()
}
