package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	pool := NewConnPool(3)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := range 5 {
		go func() {
			defer wg.Done()
			// get conn
			time.Sleep(time.Duration(rand.IntN(100)) * 100)
			conn, err := pool.Acquire(context.Background())
			if err != nil {
				fmt.Printf("acquire: %s\n", err)
				return
			}

			defer pool.Release(conn)
			fmt.Printf("%d using %d\n", i, conn.ID)
		}()
	}
	wg.Wait()
}

func NewConnPool(size int) *ConnPool {
	c := ConnPool{
		ch: make(chan Conn, size),
	}

	for i := range size {
		c.ch <- Conn{i}
	}

	return &c
}

func (p *ConnPool) Release(c Conn) {
	p.ch <- c
}

func (p *ConnPool) Acquire(ctx context.Context) (Conn, error) {
	select {
	case c := <-p.ch:
		return c, nil
	case <-ctx.Done():
		return Conn{}, ctx.Err()
	}

}

type ConnPool struct {
	ch chan Conn
}

type Conn struct {
	ID int
}
