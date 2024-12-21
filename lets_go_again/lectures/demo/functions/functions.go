package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func great() {
	fmt.Println("This is a great function!")
}

func main() {
	great()

	stack := double(32)

	backersDozen := add(double(6), 1)

	stackPlusOne := add(stack, 1)

	fmt.Println(
		"The stack in minecraft is", stack,
		"\nThe backer's dozen =", backersDozen,
		"\nAdd one to stack in minecraft and get", stackPlusOne,
	)
}
