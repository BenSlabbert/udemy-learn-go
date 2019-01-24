package main

import "fmt"

func main() {

	xInt := []int{1, 2, 3}

	fmt.Println(xInt)
	// spread out the slice to varargs of int
	sum := sumArray(xInt...)
	fmt.Println("sum: ", sum)
}

func sumArray(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}
