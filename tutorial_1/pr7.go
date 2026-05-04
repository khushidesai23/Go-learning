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

	fmt.Println("** Pointers to a function **")

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memory location of the thing1 array is %p\n", &thing1)

	var result [5]float64 = square(&thing1)
	fmt.Printf("\n The squared values are %v\n", result)
	fmt.Printf("\n The value of the thing1 array is %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64{
	var thing [5]float64
	for i:=0; i<5; i++{
		thing[i] = (*thing2)[i]*(*thing2)[i]
	}
	return thing
}

