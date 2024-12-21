package main

import "fmt"

func main() {
	fmt.Printf("Your age is %d\n", 21) // replaces %d with 21

	fmt.Printf("x is %d, y is %f\n", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159
	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius) //%d comes for DECIMAL
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi) //%f comes for FLOAT

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi) //%s for STRING

	fmt.Printf("This is a %q", figure) //%q comes for quated string

	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi) //%v comes for VAR

	fmt.Printf("figure is of type %T\n", figure) //%T comes for TYPE
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	closed := true
	fmt.Printf("File closed: %t\n", closed) //%t -> bool

	fmt.Printf("%b\n", 55) //%b -> binary form

	fmt.Printf("%08b\n", 55) //%b -> binary form with 8 bits and leading zeros

	fmt.Printf("%x\n", 100) //%x -> base 10 to base 16

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %f\n", x*y)
}
