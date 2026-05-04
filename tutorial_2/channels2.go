package main
import (
	"fmt"
	"time"
	"math/rand"
)

var MAX_PIZZA_PRICE float32 = 500
var MAX_PASTA_PRICE float32 = 300

func main(){
	var pizzaChannel = make(chan string)
	var pastaChannel = make(chan string)
	var websites = []string{"zomato.com", "swiggy.com", "dominos.com", "pizzahut.com"}

	for i:= range websites{
		go checkPizzaPrice(websites[i], pizzaChannel)
		go checkPastaPrice(websites[i], pastaChannel)
	}

	sendMessage(pizzaChannel, pastaChannel)
}

func checkPizzaPrice(website string, pizzaChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var pizza_price = rand.Float32()*500
		if pizza_price < MAX_PIZZA_PRICE{
			pizzaChannel <- website
			break
		}
	}
}

func checkPastaPrice(website string, pastaChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var pasta_price = rand.Float32()*300
		if pasta_price < MAX_PASTA_PRICE{
			pastaChannel <- website
			break
		}
	}
}

func sendMessage(pizzaChannel chan string, pastaChannel chan string){
	select{
		case website := <-pizzaChannel:
			fmt.Printf("\nText sent: Found deal for pizza on %v\n", website)
		case website := <-pastaChannel:
			fmt.Printf("\nText sent: Found deal for pasta on %v\n", website)
	}
}
