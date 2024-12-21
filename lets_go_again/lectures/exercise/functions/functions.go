//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Printf("Greetings%s!\n", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func message() string {
	return "The dew point today is at 2 degrees by celsium"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func sum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// * Write a function that returns any number
func returnEmo() int {
	return 2007
}

// * Write a function that returns any two numbers
func returnTwoBurgerNums() (int, int) {
	return 9, 6
}

//* Call every function at least once

func main() {
	greet("Cyn")

	fmt.Printf("Today's message is!%s\n", message())

	fmt.Println(sum(3, 6, 9))

	// * Add three numbers together using any combination of the existing functions.
	//   - Print the result
	a, b := returnTwoBurgerNums()
	fmt.Println(returnEmo() + a + b)
}
