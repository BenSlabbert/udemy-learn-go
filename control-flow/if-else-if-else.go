package main

import "fmt"

func main() {

	var x = 5

	if x == 10 {
		fmt.Println("x = 10")
	} else if x == 5 {
		fmt.Println("x = 5")
	} else {
		fmt.Printf("x = %d\n", x)
	}
}
