package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	/*
		a, b := 1.1, 2.2
		fmt.Println(a + b)
	*/

	data := `
	{"x": 1, "y": 2}
	{"x": 3, "y": 4}
	`
	dec := json.NewDecoder(strings.NewReader(data))
loop:
	for {
		var m map[string]int
		err := dec.Decode(&m)
		switch {
		case errors.Is(err, io.EOF):
			break loop
		case err != nil:
			fmt.Println("ERROR:", err)
			break loop
		}
		fmt.Println(m)
	}
}
