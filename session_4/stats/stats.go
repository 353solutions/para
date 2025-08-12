package main

import (
	"fmt"
	"time"
)

/*
When to use generics:
- Same function, different signature
- Generic data structures (containers)

Go's generics use "gc shape stenciling"
- Two types have the same GC shape if under the hood they are the same type
- All pointers have the same GC shape
*/

func main() {
	fmt.Println(Max([]int{2, 3, 1}))
	fmt.Println(Max([]float64{2, 3, 1}))
	fmt.Println(Max([]time.Month{time.February, time.March, time.January}))
	fmt.Println(Max[int](nil))

	freq := map[string]int{
		"the":   3923,
		"a":     3902,
		"hello": 12,
	}
	fmt.Println(MaxMap(freq))
}

// Exercise: Write a generic MaxMap that gets a map and return the key with maximal value

/*
$ dlv debug .
(dlv) funcs Max
*/

// See cmp.Ordered and built-in comparable
type Ordered interface {
	~int | ~float64 | ~string
}

func zero[T any]() T {
	var v T
	return v
}

/*
func zero[T any]() (v T) {
	return
}
*/

// T is a "type constraint", not a new type
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		/*
			var zero T
			return zero, fmt.Errorf("max of empty slice")
		*/
		return zero[T](), fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

/*
func MaxFloat64s(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

*/
