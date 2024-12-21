package main

import "fmt"

func main() {
	// x, y, z := 10, 15.5, "Gophers"
	// score := []int{10, 20, 30}
	// fmt.Printf("%d, %f, %s\n", x, y, z)
	// fmt.Printf("%v\n", score)
	// fmt.Printf("%q\n", z)
	// fmt.Printf("%v, %v, %v\n", x, y, z)
	// fmt.Printf("%T, %T", z, score)

	// const x float64 = 1.422349587101
	// fmt.Printf("%.4f", x)

	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
}
