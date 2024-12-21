package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums = [3]int{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	nums[0] = 7
	fmt.Printf("%#v\n", nums)

	//nums[5] = 100

	for i, v := range nums {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(nums); i++ {
		fmt.Println("index:", i, "value:", nums[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)
}
