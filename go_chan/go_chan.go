package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

/*
Concurrency: Dealing with several things at the same time
Parallelism: Doing with several things at the same time
*/

/* Channel semantics
- send/receive will block until opposite operation(*)
	- guarantee of delivery
- receive from a closed channel will return zero value without blocking
*/

/* Concurrency patterns
fanOut - Split work to goroutines
	- Wait to finish (sync.WaitGroup)
	- Check for error (errgroup)
	- Get return value(s)
*/

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	// This is a bug in Go < 1.22(?)
	for i := range 3 {
		// i := i
		go func() {
			fmt.Println("i:", i)
		}()
	}

	time.Sleep(100 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println("msg:", msg)

	go func() {
		for i := range 5 {
			msg := fmt.Sprintf("msg %d", i)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	/* What the above "for .. range" does

	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(msg)
	}
	*/

	msg = <-ch // ch is closed
	fmt.Printf("msg: %q\n", msg)
	// zero vs closed?
	msg, ok := <-ch
	fmt.Printf("msg: %q (ok=%v)\n", msg, ok)

	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://ibm.com/no/such/page",
	}

	start := time.Now()

	// fanOutWait()
	fanOutResult()

	duration := time.Since(start)
	fmt.Printf("%d URLs in %v\n", len(urls), duration)
}

func fanOutResult() {

}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func fanOutWait(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			urlInfo(url)
		}()
	}

	wg.Wait()
}

func urlInfo(url string) {
	//	s := time.Now()
	resp, err := http.Get(url)
	//	fmt.Printf("%q: %v\n", url, time.Since(s))

	if err != nil {
		slog.Error("info", "url", url, "error", err)
	}
	slog.Info("info", "url", url, "status", resp.Status)
}
