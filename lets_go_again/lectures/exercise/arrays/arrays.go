//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name   string
	price  int
	filled bool
}

func shoppingList(list [4]Product) {
	cost := 0
	count := 0
	last := list[0]
	for i := 0; i < len(list); i++ {
		l := list[i]
		cost += l.price
		if l.filled {
			last = l
			count++
		}
	}
	fmt.Println("The last item on the list: ", last)
	fmt.Println("The total number of items: ", count)
	fmt.Println("The total cost of the items: ", cost)
}

func main() {
	theList := [4]Product{
		{name: "Shampoo", price: 50, filled: true},
		{name: "Soap", price: 5, filled: true},
		{name: "Toothpaste", price: 10, filled: true},
	}
	shoppingList(theList)
	theList[3] = Product{name: "Suncream", price: 99, filled: true}
	shoppingList(theList)

}
