package main
import (
	"fmt"
	"errors"
)

func main(){
	var printValue string = "Hello, World! Khushi here"
	printMe(printValue)

	numerator := 11
	denominator := 0
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else if remainder == 0 {
		fmt.Printf("The result of %d divided by %d is %d with no remainder\n", numerator, denominator, result)
	} else {
		fmt.Printf("The result of %d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}