package main

import "fmt"

func main() {
	// //Exercise 1
	// strs := []string{"First", "Second", "Third"}
	// for i, v := range strs {
	// 	fmt.Println(i, ": ", v)
	// }

	// //Exercise 2
	// mySlice := []float64{1.2, 5.6}

	// //mySlice[2] = 6

	// a := 10
	// mySlice[0] = float64(a)

	// //mySlice[3] = 10.10

	// mySlice = append(mySlice, float64(a))

	// fmt.Println(mySlice)

	// //Exercise 3
	// nums := []float64{1.2, 5.6, 7.9}

	// nums = append(nums, 10.1)

	// nums = append(nums, 4.1, 5.5, 6.6)
	// fmt.Println(nums)

	// n := []float64{6.3, 9.9}
	// nums = append(nums, n...)
	// fmt.Println(nums)

	// Exercise 4
	// if len(os.Args) < 2 || len(os.Args) > 10 {
	// 	log.Fatal("Please run the program with arguments (2-10 numbers)!")
	// }
	// args := os.Args[1:]
	// sum := 0
	// product := 1
	// for ar := range args {
	// 	temp, err := strconv.Atoi(args[ar])
	// 	_ = err
	// 	sum += temp
	// 	product *= temp
	// }
	// fmt.Printf("The sum = %d\nThe product = %d\n", sum, product)

	// //Exercise 5
	// nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	// for n := range nums[2:8] {
	// 	println(nums[n+1])
	// }

	//Exercise 6
	// friends := []string{"Marry", "John", "Paul", "Diana"}
	// new_friends := make([]string, 4)
	// new_new_friends := copy(new_friends, friends)
	// friends[0] = "N"
	// new_friends[1] = "J"
	// fmt.Println(new_new_friends, new_friends, friends)

	// // Exercise 7
	// friends := []string{"Marry", "John", "Paul", "Diana"}
	// new_friends := []string{"Down"}
	// new_friends = append(new_friends, friends...)
	// fmt.Println(new_friends)

	//Exercise 8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(newYears, years[0:3]...)
	newYears = append(newYears, years[8:11]...)
	fmt.Println(newYears)
}
