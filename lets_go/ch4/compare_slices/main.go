package main

import "fmt"

func main() {
	var n []int // not initialized
	fmt.Println(n == nil)

	m := []int{} // initialized
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	//fmt.Println(a == b) // can't be done

	var eq bool = true
	//a = nil
	if len(a) != len(b) {
		eq = false
	} else {
		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break
			}
		}
	}

	fmt.Println(eq)
}
