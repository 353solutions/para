package main

import (
	"fmt"
	"io"
	"os"
)

// Exercise: Write the "wc" utility by implementing io.Writer
// Start without counting words
// -c bytes
// -l lines
// -w words

func main() {
	var wc WC
	if _, err := io.Copy(&wc, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(wc)
}

type WC struct{}
