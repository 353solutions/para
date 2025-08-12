package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

// Exercise: Write the "wc" utility by implementing io.Writer
// Start without counting words
// -c bytes
// -l lines
// -w words

// go run wc.go < wc.go
func main() {
	var wc WC
	if _, err := io.Copy(&wc, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(wc)
}

// hello there
// hello t
// here

// Write implement io.Writer
func (wc *WC) Write(p []byte) (int, error) {
	wc.bytes += len(p)
	for _, b := range p {
		if b == '\n' {
			wc.lines++
		}

		if !wc.inWord && !unicode.IsSpace(rune(b)) {
			wc.inWord = true
			wc.words++
		} else if wc.inWord && unicode.IsSpace(rune(b)) {
			wc.inWord = false
		}
	}

	return len(p), nil
}

type WC struct {
	bytes int
	lines int
	words int

	inWord bool
}
