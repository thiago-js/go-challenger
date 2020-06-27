package main

//imported the package fmt
import "fmt"

func main() {
	// we have declared an array with name "nameArray" of size 2 and data-type integer
	var nameArray [2]string

	//we assigned value Himanshu to the first element in the array nameArray
	nameArray[0] = "Himanshu"

	//we assigned value Himanshu to the second element in the array nameArray
	nameArray[1] = "Singh"

	//Printing the elements in the array by accessing it by the position
	fmt.Println("First Name:", nameArray[0])
	fmt.Println("Second Name:", nameArray[1])
}
