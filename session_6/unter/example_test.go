package main_test

import (
	"fmt"
	unter "unter"
)

func ExampleRidePrice() {
	price := unter.RidePrice(5.0, true)
	fmt.Println(price)

	// for maps: fmt.Println(maps.Equal(expected, actual))

	// Output:
	// 900
}
