package main

import (
	"fmt"
	"time"
)

func main() {
	t := NewToken()
	for i := range 4 {
		go func() {
			for range time.Tick(200 * time.Millisecond) {
				fmt.Printf("%d: %s\n", i, t.Value())
			}
		}()
	}

	time.Sleep(2 * time.Second)
}

func (t Token) Value() string {

}

// Token refreshed the token every 300ms
type Token struct {
	value string
}

// 3rd party code, you can't change it
var tokN = 0

func refreshToken() string {
	tokN++
	return fmt.Sprintf("TOK-%d", tokN)
}
