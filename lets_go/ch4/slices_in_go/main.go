package main

import (
	"fmt"
)

func main() {
	var cities []string
	fmt.Println("Cities = nil: ", cities == nil)
	fmt.Printf("Cities: %#v\n", cities)
	fmt.Println(len(cities))

	//	cities[0] = "London"

	numbers := []int{2, 3, 5, 8}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Uzi", "N"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My bff is:", myFriend)

	friends[0] = "V"
	fmt.Println("My bff is:", friends[0])

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	friends = append(friends, "J")
	fmt.Println(friends)

	var n []int
	n = numbers
	fmt.Println(n)
}
