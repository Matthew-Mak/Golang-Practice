//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	Name         string
	Owner        Member
	CheckoutTime time.Time
	ReturnTime   time.Time
	IsChecked    bool
}

type Member struct {
	Name string
}

type Library struct {
	Books   []Book
	Members []Member
}

func checkOut(bk *Book, mbr *Member) {
	bk.CheckoutTime = time.Now()
	bk.Owner = *mbr
	bk.IsChecked = Checked
}

func checkIn(bk *Book) {
	bk.ReturnTime = time.Now()
	bk.IsChecked = Unchecked
}

func printCheckedBooks(lib *Library) {
	for i := 0; i < len(lib.Books); i++ {
		if lib.Books[i].IsChecked {
			fmt.Println("Books:")
			fmt.Println(
				"	Name:", lib.Books[i].Name,
				"	\n	Owner:", lib.Books[i].Owner.Name,
				"	\n	Checkout Time:", lib.Books[i].CheckoutTime,
				"	\n	Checkin Time:", lib.Books[i].ReturnTime)
			fmt.Println("_____________________________________________________")
		} else {
			fmt.Println("[Not Returned Yet]")
			fmt.Println("_____________________________________________________")
		}
	}
}

const (
	Checked   = true
	Unchecked = false
)

func main() {
	books := make([]Book, 4)
	members := make([]Member, 3)

	books[0] = Book{Name: "Made In Abyss"}
	books[1] = Book{Name: "The Elder Scrolls"}
	books[2] = Book{Name: "Spice And Wolf"}
	books[3] = Book{Name: "The Design"}

	members[0] = Member{Name: "John"}
	members[1] = Member{Name: "Julie"}
	members[2] = Member{Name: "Kali"}

	var library Library
	library.Books = books
	library.Members = members
	printCheckedBooks(&library)

	checkOut(&books[2], &members[0])
	checkOut(&books[3], &members[2])
	printCheckedBooks(&library)

	checkIn(&books[2])
	printCheckedBooks(&library)
}
