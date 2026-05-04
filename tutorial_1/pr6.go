package main
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

func canMakeIt (e gasEngine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Printf("You can make it! You have %v miles left in the tank\n", e.milesLeft())
	}else{
		fmt.Printf("You cannot make it! You only have %v miles left in the tank\n", e.milesLeft())
	}
}

func main(){
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
}