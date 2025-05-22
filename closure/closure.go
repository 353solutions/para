package main

import (
	"fmt"
	"time"
)

func main() {
	// for i := range 5 { // 1.22
	for i := 0; i < 5; i++ {
		// i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)
}
