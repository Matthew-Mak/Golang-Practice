//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

const age = 99

func main() {
	//var age int = 0
	switch p := age; {
	case p == 0:
		fmt.Println("newborn")
	case p <= 3:
		fmt.Println("toddler")
	case p <= 12:
		fmt.Println("child")
	case p <= 17:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
