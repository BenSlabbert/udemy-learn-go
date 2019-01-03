package main

import (
	"fmt"
)

func main() {

	foo()

	fmt.Print("hello")

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println("even number found!", i)
		}
	}

	bar()

}

func foo() {
	n, _ := fmt.Println("in foo")
	fmt.Println(n)
}

func bar() {
	fmt.Println("in bar")
}
