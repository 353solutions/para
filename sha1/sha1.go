package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(FileSHA1("http.log.gz"))
	fmt.Println(FileSHA1("sha1.go"))

}

// Exercise: Decompress only if fileName ends with .gz

// cat http.log.gz| gunzip| sha1sum
// cat sha1.go| sha1sum
func FileSHA1(fileName string) (string, error) {
	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// | gunzip
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("gzip %q: %w", fileName, err)
	}
	defer r.Close()

	// | sha1sum
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("copy %q: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
