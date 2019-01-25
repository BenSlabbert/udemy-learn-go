package main

import "fmt"

func main() {

	i := returnFunc()
	fmt.Printf("%T\n", i)
	fmt.Println("the year is", i())
}

func returnFunc() func() int {
	return func() int {
		return 1984
	}
}
