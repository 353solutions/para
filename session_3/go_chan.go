package main

import (
	"context"
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
	- buffered channels of "n" has "n" free sends
- receive from a closed channel will return zero value without blocking
- send/close to a closed channel will panic
	- channel ownership
- send/receive to/from a nil channel block forever
*/

/* Concurrency patterns
fanOut - Split work to goroutines
	- Wait to finish: sync.WaitGroup)
	- Check for error: errgroup
	- Get return value: result channel (call context)
- goroutine pool
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

	// fanOutWait(urls)
	// fanOutResult(urls)
	poolDemo(urls)

	duration := time.Since(start)
	fmt.Printf("%d URLs in %v\n", len(urls), duration)

	// buffered channel to avoid goroutine leak
	ch1, ch2 := make(chan int, 1), make(chan int, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
		/*
			select {
			case ch1 <- 1:
				// OK
			case <- time.After(time.Second):
				slog.Warn("send")
			}
		*/
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	// case <-time.After(10 * time.Millisecond):
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

func poolDemo(urls []string) {
	urlCh := make(chan string)
	go infoWorker(1, urlCh)
	go infoWorker(2, urlCh)

	for _, url := range urls {
		urlCh <- url
	}
	close(urlCh) // avoid goroutine leak
	time.Sleep(time.Second)
}

func infoWorker(id int, ch chan string) {
	for url := range ch {
		fmt.Printf("%d: %q\n", id, url)
		urlInfo(url)
	}
	fmt.Printf("%d done\n", id)
}

func fanOutResult(urls []string) {
	ch := make(chan checkResult)

	for _, url := range urls {
		go func() {
			r := checkResult{url: url}
			r.status, r.err = urlCheck(url)
			ch <- r
		}()
	}
	// collect results
	for range urls {
		r := <-ch
		if r.err != nil {
			slog.Error("check", "url", r.url, "error", r.err)
			continue
		}
		slog.Info("check", "url", r.url, "status", r.status)
	}

}

type checkResult struct {
	// Call context
	url string

	status int
	err    error
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
