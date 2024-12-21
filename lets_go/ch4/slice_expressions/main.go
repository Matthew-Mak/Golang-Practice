package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	//a[start:stop]

	b := a[1:3] //includes all els in range from 1 to 3 excluding the last one: {2, 3}
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) //acceptable
	fmt.Println(s1[:3])
	fmt.Println(s1[:])
	//fmt.Println(s1[:100])//unacceptable

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)

}
