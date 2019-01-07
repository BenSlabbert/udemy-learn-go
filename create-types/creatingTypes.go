package main

import "fmt"

func main() {

	var a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	type hotdog int

	var b hotdog = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
