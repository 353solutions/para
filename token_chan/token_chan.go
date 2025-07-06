package main

import (
	"context"
	"fmt"
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
	cmd := Cmd{
		ch: make(chan string, 1), // Don't block tokenWorker on return
	}
	t.ch <- cmd
	tok := <-cmd.ch
	return tok
}

func NewToken() *Token {
	ctx, cancel := context.WithCancel(context.Background())
	t := Token{
		cancel: cancel,
		ch:     make(chan Cmd),
	}

	go tokenWorker(ctx, t.ch)

	return &t
}

func tokenWorker(ctx context.Context, ch chan Cmd) {
	tick := time.NewTicker(300 * time.Millisecond)
	tok := RefreshToken()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			tok = RefreshToken()
		case cmd := <-ch:
			cmd.ch <- tok
		}
	}
}

type Cmd struct {
	ch chan string
}

// Token refreshed the token every 300ms
type Token struct {
	ch     chan Cmd
	cancel context.CancelFunc
}

// 3rd party code, you can't change it or read from tokN
func RefreshToken() string {
	tokN++
	return fmt.Sprintf("TOK-%d", tokN)
}

var tokN = 0
