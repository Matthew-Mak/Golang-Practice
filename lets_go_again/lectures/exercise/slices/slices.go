//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(line []Part) {
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	fmt.Println("---------------------------------------")
}

func main() {
	alterLine := make([]Part, 3) // Alternative way to make a slice
	_ = alterLine
	assemblyLine1 := []Part{"Creosote Oil", "Electrite Ingot", "Silver Plate"}
	printAssemblyLine(assemblyLine1)
	assemblyLine1 = append(assemblyLine1, "Steel Rod", "Redstone Bucket")
	printAssemblyLine(assemblyLine1)
	assemblyLine1 = assemblyLine1[3:]
	printAssemblyLine(assemblyLine1)
}
