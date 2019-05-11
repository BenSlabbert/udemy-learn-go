package main

import "fmt"

func main() {

	s := "this is a String"
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)

	// e - int value
	for e := range bs {
		fmt.Println("e: ", e, " bs[e]: ", bs[e])
	}

	// k - index
	// v int value
	for k, v := range bs {
		fmt.Println("key: ", k, " value: ", v)
		fmt.Printf("at index %d we have hex %#x\n", k, v)
	}

	// UTF-8 code point
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
