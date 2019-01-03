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
	fmt.Println("in foo")
}

func bar() {
	fmt.Println("in bar")
}
