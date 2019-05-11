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

func main() {

	p1 := person{
		first: "james",
		last:  "bond",
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)

	agent := secretAgent{
		person: p1,
		ltk:    true,
	}

	fmt.Println(agent)
	fmt.Println(agent.person)
	fmt.Println(agent.person.first)
	fmt.Println(agent.first)
	fmt.Println(agent.ltk)
}
