package main
import (
	"fmt"
	"strings"
)

func main(){
	myString := "Résumé"
	indexed := myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}	  

	fmt.Printf("\n The length of the string is %v\n", len(myString))

	myString2 := []rune("Résumé")
	fmt.Printf("\n The length of the rune slice is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"Go", "is", "fun"}
	var strBuilder strings.Builder
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
		strBuilder.WriteString(" ")
	}	
	catStr := strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
