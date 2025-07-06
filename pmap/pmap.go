package main

import "fmt"

func main() {
	vals := PMap(int2str, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	fmt.Printf("%#v\n", vals) // ["1", "2", "3", "4", "5", "6", "7", "8", "9"]

	iVals := PMap(inc, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	fmt.Printf("%#v\n", iVals) // [2, 3, 4, 5, 6, 7, 8, 9, 10]

}

func inc(n int) int {
	return n + 1
}

func int2str(i int) string {
	return fmt.Sprintf("%d", i)
}

// PMap will run concurrently "fn" on every element of values.
// It will run up to "n" goroutines at the same time
// Use golang.org/x/sync/semaphore
func PMap[V any, R any](fn func(V) R, values []V, n int) []R {
	return nil // FIXME:
}
