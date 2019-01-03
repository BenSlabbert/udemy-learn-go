package main

import (
	"fmt"
)

func main() {

	var x = 42
	fmt.Printf("%T\n", x)

	var y = 'c'
	fmt.Printf("%T\n", y)

	var z = "a"
	fmt.Printf("%T\n", z)

	var a = 0.0
	fmt.Printf("%T\n", a)

	var b = true
	fmt.Printf("%T\n", b)
}
