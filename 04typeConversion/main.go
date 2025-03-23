package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" Thank for shopping with us ! ")
	fmt.Println(" Please rate our service from 1 to 5 : ")

	reader := bufio.NewReader(os.Stdin)

	input, error := reader.ReadString('\n')

	if error != nil {
		fmt.Println("something is wrong please try again Error : ", error)
	} else {
		fmt.Println("Thanks for rating us ", input)
	}

	conString, conError := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if conError != nil {
		fmt.Println("error in string conversion", conError)
	} else {
		fmt.Println("we haved added 1 extra point in your raiting and now it is : ", (conString + 1))
	}

}
