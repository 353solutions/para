package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	// _ "embed" // if you only use //go:embed directives
	"embed"
)

// $ go build -ldflags='-X main.version=0.2.0'
// Must be var (not const) and package level variable (not field in struct)
var version = "<dev>"

var options struct {
	showVersion bool
	cow         bool
}

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
	flag.BoolVar(&options.cow, "cow", false, "use cowsay API")
	flag.Usage = func() {
		prog := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "usage: %s TEXT\n", prog)
		flag.PrintDefaults()

	}
	flag.Parse()

	if options.showVersion {
		fmt.Printf("%s version %s\n", path.Base(os.Args[0]), version)
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	text := flag.Arg(0)
	if options.cow {
		cow(text)
		os.Exit(0)
	}

	width := len(text)
	fmt.Printf("  %s\n", strings.Repeat("-", width))
	fmt.Printf("< %s >\n", text)
	fmt.Printf("  %s\n", strings.Repeat("-", width))
	fmt.Println(gopher)
}

//go:embed gopher.txt
var gopher string

//go:embed *
var assets embed.FS

/*
	Usage for embed

- Embedded type can be string, []byte or embed.FS
- Static assets for web servers (js, css ...)
  - See https://pkg.go.dev/embed#example-package and http.StripPrefix
  - You can embed text/html templates but need to compile them at startup

- SQL files
*/
// Exercise: Add a --image flag that will use image from the images directory
// If not specified use gopher.txt
// Embed the images directory to the executable
func cow(text string) error {
	q := url.Values{}
	q.Add("message", text)
	q.Add("format", "text")
	url := "https://cowsay.morecode.org/say?" + q.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%q - %s", url, resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	fmt.Println()
	return err
}

/*
$ GOOS=darwin GOARCH=arm64 go build
$ go tool dist list
$ ldd gosay (Linux)
$ otool -L (OSX)
*/
