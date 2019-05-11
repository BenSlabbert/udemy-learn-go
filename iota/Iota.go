package main

import "fmt"

func main() {

	const (
		a = iota
		b
		c
	)

	// type int
	// auto increment
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
