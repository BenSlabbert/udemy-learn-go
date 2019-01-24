package main

import (
	"fmt"
)

func main() {

	// composite literal
	x := []int{1, 2, 3, 4}

	//length of slice
	fmt.Println(len(x))

	// capacity
	fmt.Println(cap(x))

	// print slice
	fmt.Println(x)

	// get value at index
	fmt.Println(x[0])

	sliceASlice(x)

	// index and value
	for i, v := range x {
		fmt.Println(i, v)
	}

	for j := 0; j < len(x); j++ {
		fmt.Println(j, x[j])
	}

	appendToSlice(x)
}

func sliceASlice(x []int) {

	//	slice a slice
	// from 1 to the end
	fmt.Println(x[1:])

	// from 1 to 3 exclusive
	fmt.Println(x[1:3])

	// from 0 to 3 exclusive
	fmt.Println(x[:3])
}

func appendToSlice(x []int) {

	x = append(x, 1, 2, 3)
	fmt.Println(x)

	y := []int{5, 6, 7, 8}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
