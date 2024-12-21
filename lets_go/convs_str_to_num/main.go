package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	s1 := "3.123"
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1 * 2)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("The i type is %T, i value is %v\n", i, i)
	fmt.Printf("The s2 type is %T, s2 value is %q\n", s2, s2)
}
