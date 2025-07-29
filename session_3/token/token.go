package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	t := NewToken()
	defer t.Close()

	for i := range 4 {
		go func() {
			for range time.Tick(200 * time.Millisecond) {
				fmt.Printf("%d: %s\n", i, t.Value())
			}
		}()
	}

	time.Sleep(2 * time.Second)
}

func (t *Token) Close() {
	t.cancel()
}

func (t *Token) Value() string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.value

}

func NewToken() *Token {
	ctx, cancel := context.WithCancel(context.Background())
	t := Token{
		value:  RefreshToken(),
		cancel: cancel,
	}

	go func() {
		// BUG: Goroutine leak
		// for range time.Tick(300 * time.Millisecond) {
		tick := time.NewTicker(300 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				tok := RefreshToken()
				// Rule of thumb: Do as less as possible inside mutex
				t.mu.Lock() // write lock
				t.value = tok
				t.mu.Unlock()
			case <-ctx.Done():
				fmt.Println("Close")
				return
			}
		}
	}()

	return &t
}

// Token refreshed the token every 300ms
type Token struct {
	mu     sync.RWMutex // single writer, multiple readers
	value  string
	cancel context.CancelFunc
}

// 3rd party code, you can't change it or read from tokN
func RefreshToken() string {
	tokN++
	return fmt.Sprintf("TOK-%d", tokN)
}

var tokN = 0
