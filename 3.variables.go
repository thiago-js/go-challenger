// Remember all programs go, needs to minimal implementation for working!

package main

import "fmt"

func main() {
	// Comments
	// In another languages this inferency not normal behavior
	// go have the inferency and strong type
	// this very awesome because this language is dynamic and another dynamic

	// Facts about variables,
	//  * := can only be used inside a function.
	//  * var is used to define global variables.
	//  * _ is a specific type of variable. Any value assigned to it like _ := 12 will be ignored.
	//  * If a variable declared in a program is not used, will throw an error.
	//  * A global variable is a variable with global scope, meaning that it is accessible throughout the program from any point

	// We declare variable in Go
	var variableName int

	// We can define multiple  variables with same dataTypes
	var variableName1, variableName2, variableName3 bool

	//Here firstName is variable of data-type string
	var firstName1 string = "Himanshu"

	// Here multiple variable are declared of same data-type
	var firstName2, lastName1 string = "Name2", "Name1"

	// We can also define variable with a initial value
	var variableName4 string = "variableName4"

	// Here firstName is variable where it is assigned a value Himanshu
	var firstName3 string = "Himanshu"

	// We can also define multiple variables with initial values
	var variableName5, variableName6 bool = false, true

	//Here multiple variable are declared of same data-type
	var firstName4, lastName2 string = "Himanshu", "Singh"

	// printLn to values
	fmt.Printf("variableName: %d  \n", variableName)
	fmt.Printf("variableName1: %t \n", variableName1)
	fmt.Printf("variableName2: %t \n", variableName2)
	fmt.Printf("variableName3: %t \n", variableName3)
	fmt.Printf("variableName4: %s \n", variableName4)
	fmt.Printf("variableName5: %t \n", variableName5)
	fmt.Printf("variableName6: %t \n", variableName6)

	fmt.Printf("firstName1: %s \n", firstName1)
	fmt.Printf("firstName2: %s \n", firstName2)
	fmt.Printf("firstName3: %s \n", firstName3)
	fmt.Printf("firstName3: %s \n", firstName4)

	fmt.Printf("lastName1: %s  \n", lastName1)
	fmt.Printf("lastName2: %s  \n", lastName2)
}
