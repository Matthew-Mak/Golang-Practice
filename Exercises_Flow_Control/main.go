package main

import (
	"fmt"
)

func main() {
	// //Eexercise 1
	// for i := 1; i <= 50; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println("The num is devisible by 7: ", i)
	// 	}
	// }

	// // Exercise 2
	// for i := 1; i <= 50; i++ {
	// 	if i%7 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println("The num is devisible by 7: ", i)
	// }

	// // Exercise 3
	// count := 0
	// for i := 1; i <= 50; i++ {
	// 	if count == 3 {
	// 		break
	// 	}
	// 	if i%7 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println("The num is devisible by 7: ", i)
	// 	count++
	// }
	// fmt.Println("Our table is broken")

	// // Exercise 4
	// for i := 1; i <= 500; i++ {
	// 	if i%5 == 0 && i%7 == 0 {
	// 		fmt.Println("It devides by 5 and 7:", i)
	// 	}
	// }

	// // Exercise 5
	// yearOfBirth := 1999
	// currentYear := time.Now().Year()
	// for i := yearOfBirth; i <= currentYear; {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Exercise 6
	age := 9
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
