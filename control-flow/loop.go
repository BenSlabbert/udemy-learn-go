package main

import "fmt"

func main() {

	// no while in golang

	for i := 0; i < 10; {

		fmt.Println(i)

		if i == 5 {
			break
		}

		i++
	}

	for i := 0; i < 10; {

		fmt.Println(i)

		if i == 4 {
			fmt.Println("found 4!")
			i++
			continue
		}

		i++
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	var a = 2
	var b = 5

	for a*b < 100 {
		a++
		b++

		fmt.Printf("a: %d, b: %d\n", a, b)
	}

	fmt.Println("done!")
}
