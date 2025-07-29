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

	// concat two slices
	s := []int{1, 2, 3, 4}
	s1 := s[:3:3] // set also capacity
	s1[0] = -1
	fmt.Println(len(s1), cap(s1))
	s2 := []int{500} //, 600}
	s3 := append(s1, s2...)
	fmt.Println(s3)
	fmt.Println(s)

	arr := [3]int{1, 2, 3} // arr is an array, arrays are passed by value
	// array type includes the size
	fmt.Println("arr:", arr)
	s = arr[:]
	fmt.Println("s:", s) // work with slice
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

	/*
		var a [10]int // array
		m := a[:]     // slice looking at a
	*/
}
