package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// _ "embed" // if you only use //go:embed directives
	"embed"
)

// $ go build -ldflags='-X main.version=0.2.0'
// Must be var (not const) and package level variable (not field in struct)
var version = "<dev>"

var options struct {
	showVersion bool
}

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
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

/* Usage for embed
- Embedded type can be string, []byte or embed.FS
- Static assets for web servers (js, css ...)
	- See https://pkg.go.dev/embed#example-package and http.StripPrefix
	- You can embed text/html templates but need to compile them at startup
- SQL files
*/
