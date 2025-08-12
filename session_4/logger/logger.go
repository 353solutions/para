package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	l := NewLogger(w)
	l.Info("All your bases are belong to us")
}

func (l Logger) Info(format string, args ...any) {
	fmt.Fprint(l.w, "INFO: ")
	fmt.Fprintf(l.w, format, args...)
	l.f.Flush()
}

func (f *fileFlusher) Flush() error {
	return f.f.Sync()
}

type fileFlusher struct {
	f *os.File
}

func (nopFlusher) Flush() error { return nil }

type nopFlusher struct{}

type flusher interface {
	Flush() error
}

func NewLogger(w io.Writer) Logger {
	l := Logger{
		w: w,
		f: nopFlusher{},
	}

	if f, ok := w.(flusher); ok {
		l.f = f
	}

	if f, ok := w.(*os.File); ok {
		l.f = &fileFlusher{f}
	}

	return l
}

type WriteFlusher interface {
	io.Writer
	Flush() error
}

type Logger struct {
	w io.Writer
	f flusher
	// w io.WriteFlusher // won't use this, limits what the logger can work with
}
