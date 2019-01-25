package main

import "fmt"

func main() {

	i := func(x int) {
		fmt.Println("func expression", x)
	}

	i(42)
	i(1984)
}
