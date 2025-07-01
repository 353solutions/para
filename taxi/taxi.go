/*
Write a function that gets an index file with names of files and sha256
signatures in the following format
0c4ccc63a912bbd6d45174251415c089522e5c0e75286794ab1f86cb8e2561fd  taxi-01.csv
f427b5880e9164ec1e6cda53aa4b2d1f1e470da973e5b51748c806ea5c57cbdf  taxi-02.csv
4e251e9e98c5cb7be8b34adfcb46cc806a4ef5ec8c95ba9aac5ff81449fc630c  taxi-03.csv
...

You should compute concurrently sha256 signatures of these files and see if
they math the ones in the index file.

  - Print the number of processed files
  - If there's a mismatch, print the offending file(s) and exit the program with
    non-zero value

$ cd /tmp
$ curl -LO  https://storage.googleapis.com/353solutions/c/data/taxi.tar
$ tar xf taxi.tar

# The index file is sha256sum.txt

Extra: Limit number of goroutines to "n" (say 4)
*/
package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	rootDir := "/tmp/taxi" // Change this to where you unpacked taxi.tar
	sigFile := path.Join(rootDir, "sha256sum.txt")
	file, err := os.Open(sigFile)

	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	start := time.Now()
	ok := true
	// unbuffered will deadlock since sigWorker is stuck sending to ch
	// and we read it only after launching all goroutines
	ch := make(chan result, len(sigs))
	sema := make(chan bool, 4)
	for name, signature := range sigs {
		sema <- ok
		fileName := path.Join(rootDir, name) + ".bz2"
		go func() {
			defer func() { <-sema }()
			sigWorker(fileName, signature, ch)
		}()
	}

	for range sigs {
		r := <-ch
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", r.fileName, r.err)
			ok = false
			continue
		}

		if !r.match {
			ok = false
			fmt.Printf("error: %s mismatch\n", r.fileName)
		}
	}

	duration := time.Since(start)
	fmt.Printf("processed %d files in %v\n", len(sigs), duration)
	if !ok {
		os.Exit(1)
	}
}

func sigWorker(fileName, signature string, ch chan<- result) {
	r := result{fileName: fileName}
	sig, err := fileSig(fileName)
	if err != nil {
		r.err = err
	} else {
		r.match = sig == signature
	}

	ch <- r
}

type result struct {
	// call context
	fileName string

	match bool
	err   error
}

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Parse signature file. Return map of path->signature
func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Line example
		// 6c6427da7893932731901035edbb9214  nasa-00.log
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			// TODO: line number
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}
