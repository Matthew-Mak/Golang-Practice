//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   bool = true
	Inactive bool = false
)

func activateItem(item *bool) {
	*item = Active
}

func deactivateItem(item *bool) {
	*item = Inactive
}

func checkout(itemList *[]bool) {
	for i := 0; i < len(*itemList); i++ {
		deactivateItem(&itemList[i])
	}
}

func main() {
	pens := false
	activateItem(&pens)
	pencils := false
	activateItem(&pencils)
	banana := false
	activateItem(&banana)
	doorHandles := false
	activateItem(&doorHandles)
	stiffCocks := false
	activateItem(&stiffCocks)

	items := make([]bool, 5)
	items[0] = pens
	items[1] = pencils
	items[2] = banana
	items[3] = doorHandles
	items[4] = stiffCocks
	for _, element := range items {
		fmt.Println(element)
	}
	fmt.Println("________________________________________")

	items[2] = Inactive
	for _, element := range items {
		fmt.Println(element)
	}
	fmt.Println("________________________________________")

	checkout(&items)
	for _, element := range items {
		fmt.Println(element)
	}
	fmt.Println("________________________________________")

}
