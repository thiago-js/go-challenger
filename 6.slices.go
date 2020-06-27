package main

//imported the package fmt
import "fmt"

func main() {

	//created a slice userdata with default value FirstName,SecondName and datatype string
	var userData = []string{"FirstName", "SecondName"}
	//print the default slice
	fmt.Println("Slice Is: ", userData)

	//replaced the element on 0th position in the slice with Himanshu
	userData[0] = "Himanshu"
	//replaced the element on 1th position in the slice with Singh
	userData[1] = "Singh"
	//prints the new slice
	fmt.Println("Slice Is: ", userData)

	//here we appends a new value "Developer" using append in the slice userData
	userData = append(userData, "Developer")
	fmt.Println("Slice Is: ", userData)

	//here we calculate the length of the slice
	fmt.Println("lenght of slice:", len(userData))

	//here we print the slice from 1st position in the slice
	fmt.Println("Slice Is: ", userData[1:])

}
