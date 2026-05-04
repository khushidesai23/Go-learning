package main
import "fmt"

func main(){
	var intSlice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice[float32](float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}


// package main
// import "fmt"

// func main(){
// 	var intSlice = []int{}
// 	fmt.Println(isEmpty[int](intSlice))

// 	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
// 	fmt.Println(isEmpty[float32](float32Slice))
// }

// func isEmpty[T any](slice []T) bool {
// 	return len(slice) == 0
// }
