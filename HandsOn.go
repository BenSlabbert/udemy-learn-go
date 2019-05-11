package main

import "fmt"

func main() {

	m := map[string]string{
		"n1": "v1",
		"n2": "v2",
		"n3": "v3",
	}

	fmt.Println(m)

	for i := range m {
		v, ok := m[i]
		fmt.Println(v, ok)
	}

	m["new"] = "newValue"

	for i := range m {
		v, ok := m[i]
		fmt.Println(v, ok)
	}

	delete(m, "new")

	for i := range m {
		v, ok := m[i]
		fmt.Println(v, ok)
	}
}
