package main
import (
	"fmt"
	"errors"
)

func main(){
	var printValue string = "Hello, World! Khushi here"
	printMe(printValue)

	numerator := 11
	denominator := 2
	result, remainder, err := intDivision(numerator, denominator)

	switch{
	case err != nil:
		fmt.Println("Error:", err)
	case remainder==0:
		fmt.Printf("The result of the integer devision is %v\n", result)
	default:
		fmt.Printf("The result of the integer devision is %v with a remainder of %v\n", result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Printf("The devision was exact")
	case 1,2:
		fmt.Printf("The devision was close")
	default:
		fmt.Printf("The devision was not close")
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