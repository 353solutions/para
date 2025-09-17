package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"slices"
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
	image       string
	listImages  bool
}

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
	flag.BoolVar(&options.cow, "cow", false, "use cowsay API")
	flag.StringVar(&options.image, "image", "", "image to use")
	flag.BoolVar(&options.listImages, "list-images", false, "list images")
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

	if options.listImages {
		dirs, err := fs.ReadDir(images, "images")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: read images directory - %s\n", err)
			os.Exit(1)

		}

		names := make([]string, 0, 1+len(dirs))
		names = append(names, "gopher")
		for _, d := range dirs {
			names = append(names, d.Name()[:len(d.Name())-4])
		}
		// Show to humans
		slices.Sort(names)
		for _, name := range names {
			fmt.Println(name)
		}
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	if options.cow && options.image != "" {
		fmt.Fprintln(os.Stderr, "error: can't use --cow with --image")
		os.Exit(1)
	}

	text := flag.Arg(0)
	if options.cow {
		cow(text)
		os.Exit(0)
	}

	var file fs.File
	if userImage() {
		var err error
		file, err = images.Open("images/" + options.image + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: load %q - %s\n", options.image, err)
			os.Exit(1)
		}
	}

	width := len(text)
	fmt.Printf("  %s\n", strings.Repeat("-", width))
	fmt.Printf("< %s >\n", text)
	fmt.Printf("  %s\n", strings.Repeat("-", width))
	if !userImage() {
		fmt.Println(gopher)
	} else {
		io.Copy(os.Stdout, file)
	}
}

func userImage() bool {
	return options.image != "" && options.image != gopherName
}

//go:embed gopher.txt
var gopher string

const gopherName = "gopher"

//go:embed images
var images embed.FS

/*
var fsImages fs.FS

func init() {
	var err error
	fsImages, err = fs.Sub(eImages, "images")
	if err != nil {
		panic(err)
	}
}
*/

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
// Extra: Add --list-images
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

$ go run . --help
$ go run . --version
$ go run . Gopher
$ go run . --cow Gopher
*/
