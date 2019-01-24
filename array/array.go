package main

import "fmt"

func main() {

	// use slices instead

	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Printf("length: %d", len(arr))
}
