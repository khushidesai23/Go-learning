//pointers tutorial
package main
import "fmt"

func main(){
	var x int = 10
	var p *int = &x
	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Address of x: %v\n", &x)
	fmt.Printf("Value of p: %v\n", p)
	fmt.Printf("Value at address p: %v\n", *p)
	*p = 20
	fmt.Printf("Value of x after change: %v\n", x)

	var p1 *int32 = new(int32)
	var i int32
	fmt.Printf("Value of p1: %v\n", p1)
	fmt.Printf("Value at address p1: %v\n", *p1)
	p1 = &i
	*p1 = 30
	fmt.Printf("Value at address p1 after change: %v\n", *p1)
	fmt.Printf("Value of i is: %v\n", i)

	fmt.Println("** Pointer to a slice **")

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Printf("Original slice: %v\n", slice)
	fmt.Printf("Copied slice: %v\n", sliceCopy)

}

