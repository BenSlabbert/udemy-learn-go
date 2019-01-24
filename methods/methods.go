package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// add this func to the secretAgent struct
// its not called a method
func (s secretAgent) speak() {
	fmt.Println("my name is", s.first, ",", s.last)
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)

	sa1.speak()
}
