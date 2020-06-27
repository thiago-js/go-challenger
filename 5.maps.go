package main

import (
	"fmt"
)

func main() {
	//Declared a map userData using the keyword `make`
	var userData = make(map[int]string)

	//prints the default map
	fmt.Println(userData)

	//adds a element at the 0th key
	userData[0] = "Himanshu"
	fmt.Println(userData)

	//adds a element at the 1th key
	userData[1] = "Singh"
	//prints the new map
	fmt.Println(userData)

	//extracts the value to a variable using the key
	var firstName = userData[0]
	fmt.Println(firstName)

	//prints the lenght of the map
	fmt.Println(len(userData))

	//delete is called to delete any value at a particular key
	delete(userData, 0)
	fmt.Println(userData)

}
