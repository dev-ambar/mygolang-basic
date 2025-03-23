package main

import (
	"bufio"
	frmt "fmt"
	"os"
)

func main() {

	frmt.Println(" Thank for shopping with us ! ")

	readr := bufio.NewReader(os.Stdin)

	frmt.Println(" Please rate our service from 1 to 5 : ")

	// comma ok // error ok syntax if anything goes wrong it commes in eroor variable
	// readString('\n') reads the input until it finds the new line character

	input, error := readr.ReadString('\n')

	// if the error is not nil then print the error
	if error != nil {
		frmt.Println("something is wrong please try again Error : ", error)
	} else {
		frmt.Println("Thanks for rating us ", input)
	}

}
