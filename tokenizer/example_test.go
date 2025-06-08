package tokenizer_test

import (
	"fmt"
	"para/tokenizer"
)

func ExampleTokenize() {
	text := "Who's on first?"
	for _, tok := range tokenizer.Tokenize(text) {
		fmt.Println(tok)
	}

	// Output:
	// first
}
