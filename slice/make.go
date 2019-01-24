package main

import "fmt"

func main() {

	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println("len: ", len(x))
	fmt.Println("cap: ", cap(x))

	x = append(x, 444)
	x = append(x, 444)
	x = append(x, 444)

	fmt.Println(x)
	fmt.Println("len: ", len(x))
	fmt.Println("cap: ", cap(x))

	// capacity doubles when initial array capacity is filled
	// new array is created with double the initial capacity
}
