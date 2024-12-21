package main

import (
	"fmt"
)

func main() {
	// var i int = 3
	// var f float64 = 3.2
	// fmt.Printf("%T, %f :: %T, %d\n", float64(i), float64(i), int(f), int(f))

	// var i = 3
	// var f = 3.2
	// var s1, s2 = "3.14", "5"

	// is := strconv.Itoa(i)
	// si2, err := strconv.Atoi(s2)
	// _ = err
	// fs := fmt.Sprintf("%f", f)
	// sf1, err := strconv.ParseFloat(s1, 64)
	// fmt.Printf("%T, %v :: %T, %v :: %T, %s :: %T, %v\n", is, is, si2, si2, fs, fs, sf1, sf1)

	// x, y := 4, 5.1

	// z := float64(x) * y
	// fmt.Println(z)

	// a := 5
	// b := 6.2 * float64(a)
	// fmt.Println(b)

	// var sm float64 = 299792458
	// var dk float64 = 149.6 * 1000000
	// fmt.Println("Time:", int((dk*1000)/sm))

	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z || int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1*5 || !(int(z) == 0)
	fmt.Println(result2)
}
