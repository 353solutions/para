package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"runtime"
	"time"
)

func main() {
	/*
		for i := range 4 { // generate 4 numbers

			if !yield(i) { // Use number
				// User ended iteration
				break
			}
		}
	*/
	r := Ints() // No code runs here
	_ = r
	for i := range Ints() {
		fmt.Println("i:", i)
		if i > 1 {
			break // Will cause yield to return false
		}
	}

	n := 0
	for t := range Tick(time.Second) {
		fmt.Println(t)
		n++
		if n > 3 {
			break
		}
	}

	fileName := "logs.json.gz"
	fmt.Println("loading", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer gz.Close()

	start := time.Now()
	/*
		logs, err := LoadLogsSlice(gz)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		count := len(logs)
	*/
	count := 0
	logs := LoadLogs(gz)
	logs = Filter(logs, IsValid)
	for range logs {
		count++
	}

	duration := time.Since(start)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc_mb := float64(mem.Alloc) / (1 << 20)
	fmt.Printf("%d logs in %v with %.2fMB\n", count, duration, alloc_mb)
}

/* Function operators
- Filter: [1 2 3] (<3): [1 2]
- Map: [1 2 3] (*7): [7 14 21]
- Reduce: [1 2 3] sum: 6

Order is important

First 10 valid logs
	logs := LoadLogs(gz)
	logs = Filter(logs, IsValid)
	logs = Head(logs, 10)

Valid logs of the first 10 (BUG?)
	logs := LoadLogs(gz)
	logs = Head(logs, 10)
	logs = Filter(logs, IsValid)

*/

func Filter[T any](s iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for v := range s {
			if !pred(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}

	}

	return fn
}

// Load slice: 59.39s 1253.66MB
// Load seq  : 57.16s    0.97MB

func LoadLogs(r io.Reader) iter.Seq[Log] {
	fn := func(yield func(Log) bool) {
		dec := json.NewDecoder(r)
		for {
			var l Log
			err := dec.Decode(&l)
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				l = Log{}
			}

			if !yield(l) {
				return
			}
		}
	}

	return fn

}

func LoadLogsSlice(r io.Reader) ([]Log, error) {
	var logs []Log
	dec := json.NewDecoder(r)
	for {
		var l Log
		err := dec.Decode(&l)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		if IsValid(l) {
			logs = append(logs, l)
		}
	}

	return logs, nil
}

type Log struct {
	Host    string    `json:"host"`
	Time    time.Time `json:"time"`
	Request string    `json:"request"`
	Status  int       `json:"status"`
	Bytes   int       `json:"bytes"`
}

func IsValid(log Log) bool {
	switch {
	case log.Host == "":
		return false
	case log.Time.IsZero():
		return false
	case log.Request == "":
		return false
	case log.Status < 100 || log.Status >= 600:
		return false
	case log.Bytes < 0:
		return false
	}

	return true
}

// IMO time.Sleep(0) is like runtime.GoSched()

func Tick(d time.Duration) iter.Seq[time.Time] {
	fn := func(yield func(time.Time) bool) {
		for {
			time.Sleep(d)
			if !yield(time.Now()) {
				return
			}
		}
	}

	return fn
}

func Ints() iter.Seq[int] {
	fn := func(yield func(int) bool) {
		n := 0
		for {
			if !yield(n) { // Pass n to user code in "for" loop
				break
			}
			n++
		}
	}

	return fn
}

func yield(i int) bool {
	fmt.Println("i:", i)
	return i > 2
}
