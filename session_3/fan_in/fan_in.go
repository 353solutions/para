package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	ch1, ch2 := producer(1), producer(2)

	for msg := range fanIn([]chan string{ch1, ch2}) {
		fmt.Println(msg)
	}
}

func fanIn(chans []chan string) chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func() {
			defer wg.Done()
			for msg := range ch {
				out <- msg
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func producer(id int) chan string {
	ch := make(chan string)

	go func() {
		for i := range 4 {
			msg := fmt.Sprintf("%d -> [%d]", id, i)
			ch <- msg
			time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}
