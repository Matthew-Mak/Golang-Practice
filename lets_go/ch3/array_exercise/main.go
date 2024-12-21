package main

import "fmt"

func main() {
	// // Exercise 1
	// var cities = [2]string{"Gobelets of Giants", "The point of no return"}
	// fmt.Println(cities)

	// var grades = [3]float64{2.5, 7.6}
	// fmt.Println(grades)

	// var taskDone = [...]bool{true, false, true}
	// fmt.Printf("%#v\n", taskDone)

	// for i, v := range cities {
	// 	println("City number", i+1, " is ", v)
	// }

	// for i := 0; i < len(grades); i++ {
	// 	fmt.Println(grades[i])
	// }

	// //Exercise 2
	// nums := [...]int{30, -1, -6, 90, -6}
	// count := 0
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i]%2 == 0 && nums[i] >= 0 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//Exercise 3
	myArray := [4]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[3] = 10.10

	fmt.Println(myArray)
}
