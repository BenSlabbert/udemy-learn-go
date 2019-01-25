package main

import "fmt"

func main() {

	innerFunc := func(sum int) {
		fmt.Println("inner func", sum)
	}

	sum := sum(innerFunc, 1, 2, 3, 4, 5)

	fmt.Println("total", sum)
}

func sum(f func(s int), x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	f(total)

	return total
}
