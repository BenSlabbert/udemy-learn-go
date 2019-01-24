package main

import "fmt"

func main() {

	foo()
	varargs("1", "2", "3")

	s := woo("penn")
	fmt.Println("s: ", s)

	str, b := mouse("author", "writer")

	fmt.Println(str, b)
}

func woo(s string) string {
	return s + " woo"
}

func foo() {
	fmt.Println("foo, no args")
}

// pass by value
// variadic func
func varargs(s ...string) {
	fmt.Println("varargs: ", s)
}

func mouse(fn, ln string) (string, bool) {
	return fmt.Sprint(fn, " ", ln, " ", "ola"), false
}
