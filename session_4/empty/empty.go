package main

import (
	"fmt"
)

/*
Rule of thumb: Don't use "any"
Exceptions:
- Printing
- Decoding/Marshaling
*/

func main() {
	var a any // Go < 1.18  interface{}

	a = 7
	fmt.Println("a:", a)
	// fmt.Println(a + 3) // won't compile

	a = "Hi"
	fmt.Println("a:", a)

	s := a.(string) // type assertion
	fmt.Println("s:", s)

	// i := a.(int) // panic
	i, ok := a.(int)
	if ok {
		fmt.Println("i:", i)
	} else {
		fmt.Println("not an int")
	}

	switch a.(type) { // type switch
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", a)
	}

	// Aside
	fmt.Println(string(0x2622)) // ☢
}
