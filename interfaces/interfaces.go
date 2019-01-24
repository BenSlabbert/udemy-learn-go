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

func (s secretAgent) speak() {
	fmt.Println("my name is", s.first, ",", s.last)
}

type human interface {
	speak()
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

	useHuman(sa1)
}

func useHuman(h human) {
	h.speak()
}
