package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int32) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int32) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int32) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	var theLot ParkingLot = ParkingLot{spaces: make([]Space, 10)}
	fmt.Println(theLot)
	theLot.occupySpace(2)
	fmt.Println(theLot)
	occupySpace(&theLot, 3)
	fmt.Println(theLot)
	theLot.vacateSpace(2)
	fmt.Println(theLot)
}
