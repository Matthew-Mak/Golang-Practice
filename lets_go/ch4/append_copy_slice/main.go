package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 42)
	fmt.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...) //ellipsis ... op
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, 8)
	nn := copy(dst, src)
	fmt.Println("dst:", dst, "src:", src, "nn:", nn)
	fmt.Println(len(dst), cap(dst))
}
