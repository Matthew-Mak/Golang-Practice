package main

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// 	goto todo // not allowed as x is declared inside goto loop
	// 	x := 5
	// todo:
	// 	fmt.Println("smth")
}
