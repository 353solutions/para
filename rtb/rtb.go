// RTB: Real Time Bidding
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// We have 50 msec to return an answer
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

// If bestBid finish in time, return the result it gave
// Otherwise, return default bid
func bidOn(ctx context.Context, url string) Bid {
	// TODO: Your code goes here
	return Bid{}
}

var defaultBid = Bid{
	AdURL: "http://adsЯus.com/default",
	Price: 3,
}

// Written by Algo team, time to completion varies
func bestBid(url string) Bid {
	// Simulate work
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: "http://adsЯus.com/ad7",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int // In ¢
}
