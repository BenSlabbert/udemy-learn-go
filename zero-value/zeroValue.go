package main

import (
	"fmt"
)

var y string

func main() {

	fmt.Println("value of y: ", y)
	fmt.Printf("%T\n", y)

	if y != "" {
		fmt.Println("y is not empty")
	} else {
		fmt.Println("y is empty")
	}

	var z int

	fmt.Println("value of z: ", z)
	fmt.Printf("%T\n", z)
}
