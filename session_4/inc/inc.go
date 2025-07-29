package main

import "fmt"

func main() {
	n := 41
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
	inc(&n)
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
	fmt.Println(n)
}

func inc(n *int) {
	*n++
	fmt.Printf("inc: n=%d, addr=%p\n", *n, n)
}

/* Rule of thumb: Start with value semantics
Go always passes by value
channels, maps are pointers
slice: contains a pointer

Must use pointers:
- If you have a lock (sync.Mutex & friends), protobuf
- marshal
- Change fields in struct
*/
