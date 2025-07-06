package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	start := time.Now()
	vals := PMap(int2str, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	fmt.Printf("%#v\n", vals) // []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(time.Since(start))

	start = time.Now()
	iVals := PMap(inc, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	fmt.Printf("%#v\n", iVals) // [int]{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(time.Since(start))

}

var nInc atomic.Int32

func inc(n int) int {
	nInc.Add(1)
	defer nInc.Add(-1)
	fmt.Println("inc:", nInc.Load())

	time.Sleep(100 * time.Millisecond)
	return n + 1
}

func int2str(i int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%d", i)
}

// PMap returns the application of "fn" on every element of values.
// It will run up to "n" goroutines at the same time.
// Use golang.org/x/sync/semaphore
func PMap[V any, R any](fn func(V) R, values []V, n int) []R {
	sema := semaphore.NewWeighted(int64(n))
	out := make([]R, len(values))

	for i, v := range values {
		sema.Acquire(context.Background(), 1)
		go func() {
			defer sema.Release(1)
			out[i] = fn(v)
		}()
	}

	sema.Acquire(context.Background(), int64(n))

	// wait
	return out
}
