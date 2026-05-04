package main
import "fmt"

func main(){
	var intArr [5]int32
	intArr[1] = 10
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}