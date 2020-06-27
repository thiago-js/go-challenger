package main

//  import one package
//  -- import "fmt" --

// import multiples packages
import (
	"errors"
	"fmt"
)

func main() {

	// using fmt package
	fmt.Print("using fmt package \n")

	// assignment error
	error := errors.New("using errors package \n")

	// using errors package
	fmt.Print(error)
}
