package main

//imported the package fmt
import "fmt"

func conditional(additionOfThreeNumbers int) {
	if additionOfThreeNumbers < 0 {
		fmt.Println("The Addition is Negative ", additionOfThreeNumbers)
	} else if additionOfThreeNumbers == 0 {
		fmt.Println("The Addition is equals to Zero ", additionOfThreeNumbers)
	} else {
		fmt.Println("The Addition is Positive ", additionOfThreeNumbers)
	}
}

func main() {
	//main function gets called when the program runs
	// declared the first variable with datatype int
	var firstNumber int = 10
	// declared the second variable without any datatype
	var secondNumber = 20
	// declared the third variable without any datatype and keyword Var
	thirdNumber := 30

	additionOfThreeNumbers := firstNumber + secondNumber + thirdNumber

	//will display the result of the addition in terminal
	fmt.Println(additionOfThreeNumbers)

	conditional(additionOfThreeNumbers)

	additionOfThreeNumbers = -75

	conditional(additionOfThreeNumbers)

	additionOfThreeNumbers = 0

	conditional(additionOfThreeNumbers)
}
