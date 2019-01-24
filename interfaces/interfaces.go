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

func (p person) speak() {
	fmt.Println("my name is", p.first, ",", p.last)
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
	useHuman(sa1.person)
}

func useHuman(h human) {
	h.speak()
}
