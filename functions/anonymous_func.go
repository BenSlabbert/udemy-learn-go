package main

import "fmt"

func main() {

	namedFoo()

	func(x int) {
		fmt.Println("anonymous func", x)
	}(1)
}

func namedFoo() {
	fmt.Println("named food")
}
