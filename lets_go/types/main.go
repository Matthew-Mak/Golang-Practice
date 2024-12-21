package main

import "fmt"

func main() {
	//RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T\n%d\n", r, r)

	//BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)

	//STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	//ARRY TYPE
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	//SLICE TYPE
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	//MAP TYPE
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	//STRUCT TYPE
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	//POINTER TYPE
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	//FUNCTION TYPE
	fmt.Printf("%T\n", f)
}
func f() {

}
