package main

import "fmt"

func main() {

	const a uint8 = 42 // typed const
	const b float64 = 42.0
	const c = "hello"

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
