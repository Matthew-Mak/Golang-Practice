package main

import "fmt"

func main() {
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{1, 2, 3}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Hoshino", "Shiruko", "Hina", "Yuuka"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	// ellipsis op
	a5 := [...]int{1, 2, 3, 5, 6, 7, 4, -33, 54}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("a5 length %d\n", len(a5))

	a6 := [...]int{
		1,
		2,
		3,
		4,
		5, // comma is mandatory
	}
	fmt.Println(a6)

}
