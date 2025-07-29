package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* Solution 1: Use mutex
	var mu sync.Mutex
	counter := 0
	*/
	// Solution: Use atomic
	counter := int64(0)

	const nGr = 10
	var wg sync.WaitGroup
	wg.Add(nGr)
	for range nGr {
		go func() {
			defer wg.Done()
			for range 1000 {
				time.Sleep(time.Microsecond)
				atomic.AddInt64(&counter, 1)
				/* Solution 1
				mu.Lock()
				counter++ // BUG: race condition. ++ is no atomic: read + modify + write
				mu.Unlock()
				*/
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}

// go run -race counter.go
// go build & go test also support -race
// Best practice: Run tests with -race
