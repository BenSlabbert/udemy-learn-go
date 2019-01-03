package main

import (
	"fmt"
)

var GLOBAL = 10

func main() {

	// short declaration
	// declare and assign value
	x := 42
	fmt.Println(x)

	// print address of the variable (pointers :D)
	fmt.Println(&x)

	i := &x
	fmt.Println(*i)

	declaring()
}

func declaring() {
	var y int16 = 5
	fmt.Println(y)
	useType(y)
}

func useType(p ...int16) {
	fmt.Println(p[0])
	fmt.Println(GLOBAL)
}
