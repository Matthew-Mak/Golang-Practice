package main

import "fmt"

func main() {
	const a float64 = 5.1 //Typed constant

	const b = 6 //Untyped constant

	const c float64 = a * b
	const str = "Hello " + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	var i int = b     //b changes to int
	var j float64 = b // var j float64 = float64(b)
	var p byte = b

	fmt.Println(i, j, p)

}
