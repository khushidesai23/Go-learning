package main

import (
	"fmt"
	"time"
)

func main(){
	n := 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeloop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeloop(testSlice2, n))
}

func timeloop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}