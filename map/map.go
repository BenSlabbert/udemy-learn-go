package main

import "fmt"

func main() {

	m := map[string]int{
		"james": 32,
		"money": 26,
		"penny": 67,
	}

	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["no key"]) // returns 0

	v, ok := m["no key"]

	fmt.Println(v, ok) // check if value exists

	// comma ok idiom
	if v, ok := m["no key"]; ok {
		fmt.Println("value exists: ", v)
	} else {
		fmt.Println("value does not exists: ", v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	addToMap(m)
	deleteFromMap(m)
}

func deleteFromMap(m map[string]int) {

	fmt.Println(m)
	delete(m, "james")

	for k, v := range m {
		fmt.Println(k, v)
	}

	// no errors thrown for no key in map
	delete(m, "no key")

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func addToMap(m map[string]int) {

	fmt.Println(m)
	m["added"] = 34
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
