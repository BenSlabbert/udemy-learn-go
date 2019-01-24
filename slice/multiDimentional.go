package main

import "fmt"

func main() {

	jb := []string{"james", "bond", "money penny"}
	fmt.Println(jb)

	flavors := []string{"vanilla", "strawberry", "chocolate"}
	fmt.Println(flavors)

	multi := [][]string{jb, flavors}

	fmt.Println(multi)

	fmt.Println(multi[0])
	fmt.Println(multi[0][0])
	fmt.Println(multi[1])
	fmt.Println(multi[1][0])
}
