package main

import "fmt"

func main() {

	defer fooDefer("1")
	defer fooDefer("2")
	barDefer()
}

func fooDefer(s string) {
	fmt.Println("fooDefer: ", s)
}

func barDefer() {
	fmt.Println("barDefer")
}
