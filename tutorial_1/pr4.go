package main
import "fmt"

func main(){
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length of the slice is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length of the slice is %v and the capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"Khushi": 23,
		"Priyam": 27,
	}
	fmt.Println(myMap2["Priyam"])
	fmt.Println(myMap2["Vismay"])

	var age, ok = myMap2["Vismay"]
	delete(myMap2, "Sarah")
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid name\n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr{
		fmt.Printf("The index is %v and the value is %v\n", i, v)
	}

	var i int = 0
	for{
		if i>=10 {
			break
		}
		fmt.Println(i)
		i = i+1
	}

	for i:=0; i<10; i++{
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v is odd\n", i)
	}
}