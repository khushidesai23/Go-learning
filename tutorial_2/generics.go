package main
import "fmt"

func main(){
	var intSlice = []int{}
	fmt.Println(isEmpty[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(isEmpty[float32](float32Slice))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

