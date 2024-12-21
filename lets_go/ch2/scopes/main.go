package main

import (
	"fmt"

	//import "fmt" //error

	f "fmt"
)

// permitted in go

const done = false // package scoped

func main() { // package scoped
	var task = "Running" // block scoped or local

	const done = true // local scoped
	f.Println("done in main = ", done)
	f1()
	fmt.Println(task, done)
	f.Println("bye bye (see line:7)")
}

func f1() {
	f.Println("done in f1 = ", done)
}
