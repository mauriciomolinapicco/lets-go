package main

import "fmt"
import "errors"
	
func main() {
	var printValue string = "jeluouu"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 0
	var divisionResult, divisionRemainder, err = intDivision(numerator, denominator)
	// PRINTF for formatted strings
	if err != nil {
		fmt.Printf(err.Error())
	} else if divisionRemainder == 0 {
		fmt.Printf("The result of the division is %v and there is no remainder", divisionResult)
	} else {
		fmt.Printf("The result of the division is %v and the remainder is %v", divisionResult, divisionRemainder)
	}
}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error //error is a native type. Default value is nil

	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}