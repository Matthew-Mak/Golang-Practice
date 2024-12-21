//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func areaOfRectangle(figure Rectangle) int {
	return figure.Length * figure.Width
}

func perimeterOfRectangle(figure Rectangle) int {
	return 2 * (figure.Length + figure.Width)
}

func main() {
	rect := Rectangle{Length: 5, Width: 4}
	fmt.Println("The area of the rectangle = ", areaOfRectangle(rect))
	fmt.Println("The perimeter of the rectangle = ", perimeterOfRectangle(rect))
}
