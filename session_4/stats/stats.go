package main

import (
	"cmp"
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
	fmt.Println(MaxMap(freq)) // the <nil>
	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(m.At(3, 2))
}

// Limitation: Can't have generic method arguments
func (m *Matrix[T]) At(row, col int) (T, error) {
	i := row*m.Cols + col
	if i >= len(m.data) {
		var zero T
		return zero, fmt.Errorf("%d/%d out of range for %d/%d", row, col, m.Rows, m.Cols)
	}

	return m.data[i], nil
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("%d/%d - bad dimension", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,

		data: make([]T, rows*cols),
	}

	return &m, nil
}

type Matrix[T Number] struct {
	Rows int
	Cols int

	data []T
}

type Number interface {
	~int | ~float64
}

type Handler interface {
	User | Event
	GetID() int
}

// func Handle[T User | Event](v T) {
func Handle[T Handler](v T) {
	// fmt.Println(v.ID) // Compile error: can't access field structs
	fmt.Println(v.GetID())

	// In the case of this implementation, use interfaces without generics
}

// Python uses duck typing

func (u User) GetID() int {
	return u.ID
}

type User struct {
	ID int
}

func (e Event) GetID() int {
	return e.ID
}

type Event struct {
	Name string
	ID   int
}

// type User = users.User

// Exercise: Write a generic MaxMap that gets a map and returns the key with maximal value.
// Return error on empty map.
func MaxMap[K comparable, V cmp.Ordered](m map[K]V) (K, error) {
	if len(m) == 0 {
		return zero[K](), fmt.Errorf("max on empty map")
	}

	var maxK K
	var maxV V
	first := true

	for k, v := range m {
		if first || v > maxV {
			maxK, maxV = k, v
			first = false
		}
	}

	return maxK, nil
}

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
