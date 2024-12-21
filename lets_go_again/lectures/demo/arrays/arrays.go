package main

import "fmt"

type Room struct {
	name     string
	cleanded bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleanded {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "needs cleaning")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Werehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[2].cleanded = true
	rooms[3].cleanded = true

	checkCleanliness(rooms)
}
