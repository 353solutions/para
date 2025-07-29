package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Business requirement:
	// If there's no count set it to 1
	// If there's a count, it must be > 0
	// data := []byte(`{"image": "ubuntu"}`)
	data := []byte(`{"image": "ubuntu", "count": 10}`)

	// var req StartVM
	// Solution 3: Default values
	req := StartVM{
		Count: 1,
	}
	if err := json.Unmarshal(data, &req); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%#v\n", req)
	/* Solution 2: map[string]any (mapstructures)
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	// fmt.Printf("%#v\n", req)
	fmt.Printf("%#v\n", m)

	c, ok := m["count"]
	if !ok {
		fmt.Println("ERROR: no count")
		return
	}
	// TODO: Use mapstructure instead
	count, ok := c.(float64) // type assertion
	if !ok {
		fmt.Printf("ERROR: bad count type %#v\n", c)
		return
	}
	fmt.Println("count:", count)
	// TODO: convert count to int, fail on fraction
	*/
}

type StartVM struct {
	Image string
	Count int
	// Solution 1: Use pointers for non-nullable types
	// Count *int
}
