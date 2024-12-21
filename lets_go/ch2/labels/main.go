package main

import "fmt"

func main() {
	outer := 19
	_ = outer
	people := [5]string{"Kiui", "Habolck", "Jiruo", "Riko", "Reg"}
	friends := [2]string{"Shegi", "Kiui"}

outer: // no conflict between outer lable and outer variable
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a red whistle %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}
