package main

import "fmt"

func main() {
	var nums []int
	curCap := 0
	for i := range 10_000 {
		// nums = append(nums, i)
		nums = appendInt(nums, i)
		if c := cap(nums); c != curCap {
			fmt.Println(curCap, "->", c)
			curCap = c
		}
	}
	fmt.Println(nums[:10])
}

func appendInt(s []int, n int) []int {
	if len(s) == cap(s) { // need to re-allocate & copy
		size := (2 * cap(s)) + 1
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}
	// s[len(s)+1] // panic
	s = s[:len(s)+1]
	s[len(s)-1] = n
	return s
}
