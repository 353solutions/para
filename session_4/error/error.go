package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := OpenFile("error.go"); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("OK")
	// Use errors.Is and errors.As

}

func OpenFile(path string) (*os.File, error) {
	//var err *OSError // BUG: won't be nil error since the interface tab field is not nil
	/*
		err = iface {
			data: nil,
			tab: &OSError,
		}
	*/
	var err error
	// err = iface{nil, nil}

	// TODO
	return nil, err
	// return nil, &OSError{path}
}

func (o *OSError) Error() string {
	return o.Path
}

type OSError struct {
	Path string
}
