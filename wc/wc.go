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

// Write implement io.Writer
func (wc *WC) Write(p []byte) (int, error) {
	wc.bytes += len(p)
	inWord := false
	for _, b := range p {
		if b == '\n' {
			wc.lines++
		}

		if !inWord && !unicode.IsSpace(rune(b)) {
			inWord = true
			wc.words++
		} else if inWord && unicode.IsSpace(rune(b)) {
			inWord = false
		}
	}

	return len(p), nil
}

type WC struct {
	bytes int
	lines int
	words int
}
