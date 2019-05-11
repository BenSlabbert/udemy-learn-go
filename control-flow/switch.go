package main

import "fmt"

func main() {

	// no default fallthrough, must be specified
	x := 4

	switch x {

	case 1, 4:
		fmt.Printf("x = %d\n", x)
		fallthrough

	case 10:
		fmt.Println("x = 10")

	case 11:
		fmt.Println("x = 11")

	default:
		fmt.Printf("x = %d\n", x)

	}
}
