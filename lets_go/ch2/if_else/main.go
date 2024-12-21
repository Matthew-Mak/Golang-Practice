package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive!")
	}
	//_ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

}
