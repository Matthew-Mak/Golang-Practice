package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	var name = "Dan"
	// fmt.Println("Your name is: ", name)
	_ = name

	s := "Learning golang!"
	fmt.Println(s)

	name = "Andrei"

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j //variables swap

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
