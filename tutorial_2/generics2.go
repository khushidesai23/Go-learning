package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type contactInfo struct{
	Name string
	Email string
}

type purchaseInfo struct{
	Name string
	Price float32
	Amount int
}

func main(){
	var contacts []contactInfo = loadJSON[contactInfo]("./tutorial_2/contacts.json")
	fmt.Printf("Contacts: %v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./tutorial_2/purchases.json")
	fmt.Printf("Purchases: %v\n", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T{
	data, _ := ioutil.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}

